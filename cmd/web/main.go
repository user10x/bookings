package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/pkg/driver"
	"github.com/nickhalden/mynicceprogram/pkg/handlers"
	"github.com/nickhalden/mynicceprogram/pkg/helpers"
	"github.com/nickhalden/mynicceprogram/pkg/render"
	"log"
	"net/http"
	"os"
	"time"
)

var socketAddr string
var infoLog *log.Logger
var errorLog *log.Logger
var session *scs.SessionManager

// main part of the application
func main() {

	var app config.AppConfig
	// get the connection and close db connection after main is done

	session = scs.New()

	sessionManager := session
	sessionManager.Lifetime = 24 * time.Hour
	app.Session = sessionManager

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
		log.Fatalf("error connecting to the database %v", err)
	}

	log.Println("Connected to database")

	// Repository pattern
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	mux := chi.NewRouter()

	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(NoSurf)
	mux.Use(middleware.Logger)
	mux.Use(middlewareCustom) //custom middleware
	//app.Session.LoadAndSave(mux)

	mux.Get("/health", handlers.Repo.Health)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/make-registration", handlers.Repo.MakeRegistration)
	mux.Post("/make-registration", handlers.Repo.PostRegistration)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)

	helpers.NewHelpers(&app)

	socketAddr = ":8000"
	log.Println("starting the server on", socketAddr)

	http.ListenAndServe(socketAddr, sessionManager.LoadAndSave(mux))
	//mux.Use(LoadSession)

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
