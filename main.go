package main

import (
	"language-backend/database"
	"language-backend/endpoints"
	"net/http"

	_ "modernc.org/sqlite"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // TODO fix domain
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	endpointHandler := endpoints.NewEndpointHandler(db)

	mux.HandleFunc("/verb/random", endpointHandler.VerbRandom)
	mux.HandleFunc("/verb/list", endpointHandler.VerbList)
	mux.HandleFunc("/verb/{verb}", endpointHandler.VerbDetails)

	handler := enableCORS(mux)
	http.ListenAndServe(":8080", handler)
}
