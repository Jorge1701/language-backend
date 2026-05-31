package endpoints

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	max_verb_list = 100
)

func (h *Handler) VerbList(w http.ResponseWriter, r *http.Request) {
	amountParam := r.URL.Query().Get("amount")
	if amountParam == "" {
		response400(w, r, "'amount' is required")
		return
	}

	amountValue, err := strconv.Atoi(amountParam)
	if err != nil {
		response400(w, r, "'amount' should be a number")
		return
	}

	if amountValue > max_verb_list {
		response400(w, r, fmt.Sprintf("'amount' max is %d", max_verb_list))
		return
	}

	verbs, err := h.db.ListVerbs(amountValue)
	if err != nil {
		response500(w, r, err)
		return
	}

	response200(w, r, verbs)
}
