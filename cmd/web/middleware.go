package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

func middlewareCustom(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

// todo: look and watch and use
func NoSurf(next http.Handler) http.Handler {
	csfr := nosurf.New(next)
	csfr.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return csfr
}

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
