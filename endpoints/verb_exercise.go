package endpoints

import (
	"errors"
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
	tenseParam := model.Tense(strings.TrimSpace(strings.ToUpper(r.URL.Query().Get("tense"))))
	if !tenseParam.IsValid() {
		response400(w, r, "Invalid param 'tense'")
		return
	}
	pronounParam := model.Pronoun(strings.TrimSpace(strings.ToUpper(r.URL.Query().Get("pronoun"))))
	if !pronounParam.IsValid() {
		response400(w, r, "Invalid param 'pronoun'")
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
	verbExerciseInfo := model.VerbExerciseInfo{
		Tense:       tenseParam,
		Pronoun:     pronoun,
		Verb:        verbParam,
		Conjugation: conjugation,
		Example:     h.findExample(pronoun, conjugation),
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

func (h *Handler) findExample(pronoun, conjugation string) *model.ExerciseExample {
	examples, err := h.db.FindVerbExample(pronoun, conjugation)
	if err != nil {
		return nil
	}

	curatedExamples := []model.ExerciseExample{}

	for _, example := range examples {
		parts := strings.Split(strings.ToLower(strings.TrimSpace(example.Pt)), " ")

	parts_loop:
		for i, part := range parts {
			if part == pronoun {
				if len(parts) > i+1 && parts[i+1] == conjugation {
					curatedExamples = append(curatedExamples, model.ExerciseExample{
						Pt: example.Pt,
						Es: example.Es,
					})
				}

				break parts_loop
			}
		}
	}

	if len(curatedExamples) == 0 {
		return nil
	}

	return &curatedExamples[rand.Intn(len(curatedExamples))]
}

var (
	tps = []string{"você", "ele", "ela"}
	tpp = []string{"vocês", "eles", "elas"}
)

func getPronounText(pronoun model.Pronoun) string {
	switch pronoun {
	case model.FirstPersonalSingular:
		return "eu"
	case model.SecondPersonalSingular:
		return "tu"
	case model.ThirdPersonalSingular:
		return tps[rand.Intn(3)]
	case model.FirstPersonalPlural:
		return "nós"
	case model.SecondPersonalPlural:
		return "vós"
	case model.ThirdPersonalPlural:
		return tpp[rand.Intn(3)]
	}
	return "-"
}
