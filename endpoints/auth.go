package endpoints

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

const (
	claimKey = "claim"
)

func (h *Handler) AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			response401(w, r, "Unauthorized")
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			response401(w, r, "Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), claimKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
