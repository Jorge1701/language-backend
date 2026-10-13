package endpoints

import (
	"language-backend/generator"
	"net/http"
)

func (h *Handler) Generator(w http.ResponseWriter, r *http.Request) {
	res, err := h.generator.Conjugate("falar", generator.PresenteDoIndicativo, generator.Eu)
	if err != nil {
		response500(w, r, err)
		return
	}

	response200(w, r, res)
}
