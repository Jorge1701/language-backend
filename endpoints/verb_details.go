package endpoints

import (
	"errors"
	"fmt"
	"language-backend/database"
	"net/http"
	"strings"
)

func (h *Handler) VerbDetails(w http.ResponseWriter, r *http.Request) {
	verbParam := strings.TrimSpace(strings.ToLower(r.PathValue("verb")))
	if verbParam == "" {
		response400(w, r, "'verb' is required")
		return
	}

	verbRow, err := h.db.FindVerb(verbParam)
	if err != nil {
		if errors.Is(err, database.ErrNotFoundInDB) {
			response404(w, r, fmt.Sprintf("Verb '%s' not found", verbParam))
		} else {
			response500(w, r, err)
		}
		return
	}

	response200(w, r, mapVerbRow(verbRow))
}
