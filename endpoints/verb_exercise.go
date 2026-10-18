package endpoints

import (
	"errors"
	"fmt"
	"language-backend/database"
	"language-backend/model"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrConjugationUnavailable = errors.New("Conjugation unavailable")
)

func (h *Handler) VerbExercise(w http.ResponseWriter, r *http.Request) {
	tenseParam := model.Tense(strings.TrimSpace(r.URL.Query().Get("tense")))
	if !tenseParam.IsValid() {
		response400(w, r, fmt.Sprintf("Invalid param 'tense': %s", tenseParam))
		return
	}
	pronounParam := model.Pronoun(strings.TrimSpace(r.URL.Query().Get("pronoun")))
	if !pronounParam.IsValid() {
		response400(w, r, fmt.Sprintf("Invalid param 'pronoun': %s", pronounParam))
		return
	}
	verbParam := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("verb")))
	randomLimitParamRaw := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("random_limit")))
	if verbParam == "" && randomLimitParamRaw == "" {
		response400(w, r, "Query param 'verb' or 'random_limit' is required")
		return
	}

	var err error
	var randomLimitParam = 0
	if verbParam == "" && randomLimitParamRaw != "" {
		randomLimitParam, err = strconv.Atoi(randomLimitParamRaw)
		if err != nil || randomLimitParam <= 0 {
			response400(w, r, "Query param 'random_limit' should be a positive number")
			return
		}
	}

	if verbParam == "" && randomLimitParam > 0 {
		verbRow, err := h.db.FindRandomVerb(randomLimitParam)
		if err != nil {
			handleError(w, r, err)
			return
		}
		verbParam = verbRow.Infinitive
	}

	conjugation, err := h.db.FindConjugation(tenseParam, pronounParam, verbParam)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if strings.TrimSpace(conjugation) == "" {
		response404(w, r, "Could not find conjugation")
		return
	}

	pronoun := getPronounText(pronounParam)
	example, err := h.findExample(tenseParam, pronoun, verbParam)
	if err != nil {
		response500(w, r, err)
		return
	}

	verbExerciseInfo := model.VerbExerciseInfo{
		Tense:       tenseParam,
		Pronoun:     pronoun,
		Verb:        verbParam,
		Conjugation: conjugation,
		Example:     example,
	}

	response200(w, r, verbExerciseInfo)
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, database.ErrNotFoundInDB) {
		response404(w, r, "Verb not found")
	} else {
		response500(w, r, err)
	}
}

func (h *Handler) findExample(tense model.Tense, pronoun, verb string) (*model.ExerciseExample, error) {
	examples, err := h.db.FindVerbExample(tense, pronoun, verb)
	if err != nil {
		return nil, err
	}

	if len(examples) == 0 {
		return nil, nil
	}

	example := examples[rand.Intn(len(examples))]
	return &model.ExerciseExample{
		Pt: example.Pt,
		Es: example.Es,
	}, nil
}

var (
	voceEleEla    = []string{"você", "ele", "ela"}
	vocesElesElas = []string{"vocês", "eles", "elas"}
)

func getPronounText(pronoun model.Pronoun) string {
	switch pronoun {
	case model.Eu:
		return "eu"
	case model.Tu:
		return "tu"
	case model.VoceEleEla:
		return voceEleEla[rand.Intn(3)]
	case model.Nos:
		return "nós"
	case model.Vos:
		return "vós"
	case model.VocesElesElas:
		return vocesElesElas[rand.Intn(3)]
	}
	return "-"
}
