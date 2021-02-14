package handler

import (
	"net/http"
)

// About returns about content
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Calling About route"))
}

// Home returns home route
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Calling Home route"))
}
