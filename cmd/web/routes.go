package main

//import (
//	"github.com/go-chi/chi"
//	"github.com/go-chi/chi/middleware"
//	"github.com/nickhalden/mynicceprogram/pkg/config"
//	"github.com/nickhalden/mynicceprogram/pkg/handlers"
//	"net/http"
//	"time"
//)
//
//func routes(a *config.AppConfig) {
//	r := chi.NewRouter()
//	r.Use(middleware.Timeout(60 * time.Second))
//
//	r.Use(middleware.Logger)
//	r.Use(middlewareCustom) //custom middleware
//
//	r.Get("/about", handlers.About)
//	r.Get("/home", handlers.Home)
//
//	socketAddr = ":8000"
//	http.ListenAndServe(socketAddr, r)
//
//}
