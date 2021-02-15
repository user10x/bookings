package main

import (
	"fmt"
	"github.com/nickhalden/mynicceprogram/pkg/handler"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var socketAddr string

// main part of the application
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Logger)
	r.Use(middlewareCustom) //custom middleware
	r.Get("/about", handler.About)
	r.Get("/home", handler.About)

	socketAddr = ":8000"
	http.ListenAndServe(socketAddr, r)

	fmt.Println("something out here")

}
