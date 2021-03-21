package main

import (
	"fmt"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/pkg/driver"
	"github.com/nickhalden/mynicceprogram/pkg/handlers"
	"github.com/nickhalden/mynicceprogram/pkg/helpers"
	"github.com/nickhalden/mynicceprogram/pkg/render"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var socketAddr string
var infoLog *log.Logger
var errorLog *log.Logger

// main part of the application
func main() {

	var app config.AppConfig
	// get the connection and close db connection after main is done

	// app wide variables
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	render.NewRenderer(&app)

	// connect to the database
	db, err := run()
	defer db.SQL.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	// Repository pattern
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Logger)
	r.Use(middlewareCustom) //custom middleware

	r.Get("/health", handlers.Repo.Health)
	r.Get("/about", handlers.Repo.About)
	r.Get("/home", handlers.Repo.Home)

	helpers.NewHelpers(&app)

	socketAddr = ":8000"
	log.Println("starting the server on", socketAddr)
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
