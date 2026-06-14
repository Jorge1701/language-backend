package endpoints

import (
	"net/http"
)

type MeRes struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(claimKey).(*Claims)
	if claims == nil {
		response401(w, r, "Unauthorized")
		return
	} else {
		response200(w, r, MeRes{
			UserId:   claims.UserID,
			Username: claims.Username,
		})
	}
}
