package main

import (
	"language-backend/database"
	"language-backend/endpoints"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func headersMiddleware(isProd bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if isProd {
			w.Header().Set("Access-Control-Allow-Origin", "https://portugues.rosasjorge.xyz")
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

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

	endpointHandler := endpoints.NewEndpointHandler(db)

	mux.HandleFunc("/login", endpointHandler.Login)
	mux.HandleFunc("/register", endpointHandler.Register)

	mux.HandleFunc("/verb/exercise", endpointHandler.VerbExercise)
	mux.HandleFunc("/verb/random", endpointHandler.VerbRandom)
	mux.HandleFunc("/verb/list", endpointHandler.VerbList)
	mux.HandleFunc("/verb/{verb}", endpointHandler.VerbDetails)

	isProd := os.Getenv("APP_ENV") == "prod"
	handler := headersMiddleware(isProd, mux)

	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), handler)
	if err != nil {
		panic(err)
	}
}
