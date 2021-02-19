package main

import (
	"fmt"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/pkg/driver"
	"github.com/nickhalden/mynicceprogram/pkg/handlers"
	"github.com/nickhalden/mynicceprogram/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var socketAddr string

// main part of the application
func main() {

	var app config.AppConfig
	// get the connection and close db connection after main is done
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Connected to database")
	defer db.SQL.Close()

	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Logger)
	r.Use(middlewareCustom) //custom middleware

	r.Get("/health", handlers.Repo.Health)
	r.Get("/about", handlers.Repo.About)
	r.Get("/home", handlers.Repo.Home)

	socketAddr = ":8000"
	http.ListenAndServe(socketAddr, r)

	fmt.Println("something out here")

}

func run() (*driver.DB, error) {
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=visitors user=nipun.chawla password=")
	if err != nil {
		log.Fatal("cannot connect to the database!  Dying...")
		return nil, err
	}

	return db, err

}
