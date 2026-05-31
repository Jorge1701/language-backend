package endpoints

import (
	"net/http"
	"strconv"
)

func (h *Handler) VerbRandom(w http.ResponseWriter, r *http.Request) {
	topSearchedParam := r.URL.Query().Get("top_searched")
	if topSearchedParam == "" {
		response400(w, r, "'top_searched' param is required")
		return
	}

	topSearch, err := strconv.Atoi(topSearchedParam)
	if err != nil {
		response400(w, r, "'top_searched' should be a number")
		return
	}

	verbRow, err := h.db.FindRandomVerb(topSearch)

	if err != nil {
		response500(w, r, err)
		return
	}

	response200(w, r, mapVerbRow(verbRow))
}
