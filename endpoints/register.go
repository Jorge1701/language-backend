package endpoints

import (
	"encoding/json"
	"errors"
	"language-backend/database"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	registerLimiter = rate.NewLimiter(rate.Every(time.Second*3), 1)
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if !registerLimiter.Allow() {
		response429(w, r, "Too many requests.")
		return
	}

	var req RegisterReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response400(w, r, "Invalid body, provide username and password.")
		return
	}

	if len(req.Username) < 5 || len(req.Username) > 50 {
		response400(w, r, "username must be between 3 and 50 characters.")
		return
	}

	if len(req.Password) < 8 {
		response400(w, r, "password must be at least 8 characters.")
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
