package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
	"strings"
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

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ipAddress := req.RemoteAddr
		fwdAddress := req.Header.Get("X-Forwarded-For") // capitalisation doesn't matter
		if fwdAddress != "" {
			// Got X-Forwarded-For
			ipAddress = fwdAddress // If it's a single IP, then awesome!

			// If we got an array... grab the first IP
			ips := strings.Split(fwdAddress, ", ")
			if len(ips) > 1 {
				ipAddress = ips[0]
			}
		}
		log.Println("Got connection from ", ipAddress)
		h.ServeHTTP(rw, req)
	})
}
