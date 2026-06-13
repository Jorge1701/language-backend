package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func response200(w http.ResponseWriter, r *http.Request, result any) {
	fmt.Printf("200 - %s\n%v\n", r.URL, result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func response400(w http.ResponseWriter, r *http.Request, message string) {
	fmt.Printf("400 - %s - %s\n", r.URL, message)

	errorResponse(w, message, http.StatusBadRequest)
}

func response404(w http.ResponseWriter, r *http.Request, message string) {
	fmt.Printf("404 - %s - %s\n", r.URL, message)

	errorResponse(w, message, http.StatusNotFound)
}

func response500(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("500 - %s - %v\n", r.URL, err)

	errorResponse(w, "Internal server error", http.StatusInternalServerError)
}

func errorResponse(w http.ResponseWriter, message string, status int) {
	http.Error(w, fmt.Sprintf(`{"error": "%s"}`, message), status)
}
