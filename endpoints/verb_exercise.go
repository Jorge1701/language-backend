package endpoints

import (
	"errors"
	"language-backend/database"
	"language-backend/model"
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

	verbExerciseInfo := model.VerbExerciseInfo{
		Tense:       tenseParam,
		Pronoun:     pronounParam,
		Verb:        verbParam,
		Conjugation: conjugation,
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
