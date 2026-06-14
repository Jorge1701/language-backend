package endpoints

import (
	"encoding/json"
	"errors"
	"language-backend/database"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response500(w, r, err)
		return
	}

	user, err := h.db.FindUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, database.ErrNotFoundInDB) {
			response401(w, r, "Invalid credentials.")
		} else {
			response500(w, r, err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		response401(w, r, "Invalid credentials.")
		return
	}

	claims := &Claims{
		UserID:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response500(w, r, err)
		return
	}

	response200(w, r, map[string]string{"token": signed})
}
