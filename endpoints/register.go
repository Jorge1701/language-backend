package endpoints

import (
	"encoding/json"
	"errors"
	"language-backend/database"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response400(w, r, "Invalid body, provide username and password.")
		return
	}

	_, err = h.db.FindUserByUsername(req.Username)
	if !errors.Is(err, database.ErrNotFoundInDB) {
		response409(w, r, "Username already taken.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response500(w, r, err)
		return
	}

	err = h.db.CreateNewUser(req.Username, string(hash))
	if err != nil {
		response500(w, r, err)
		return
	}

	response(w, r, http.StatusCreated)
}
