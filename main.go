package main

import (
	"language-backend/database"
	"language-backend/endpoints"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOWED_ORIGIN"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	endpointHandler := endpoints.NewEndpointHandler(
		db,
		os.Getenv("ENV") == "prod",
	)

	mux.HandleFunc("POST /login", endpointHandler.Login)
	mux.HandleFunc("POST /register", endpointHandler.Register)
	mux.HandleFunc("POST /logout", endpointHandler.Logout)

	mux.Handle("GET /me", endpointHandler.AuthRequired(
		http.HandlerFunc(endpointHandler.Me),
	))
	mux.Handle("GET /verb/exercise", endpointHandler.AuthRequired(
		http.HandlerFunc(endpointHandler.VerbExercise),
	))
	mux.HandleFunc("GET /verb/random", endpointHandler.VerbRandom)
	mux.HandleFunc("GET /verb/list", endpointHandler.VerbList)
	mux.HandleFunc("GET /verb/{verb}", endpointHandler.VerbDetails)

	handler := headersMiddleware(mux)

	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handler)
	if err != nil {
		panic(err)
	}
}
