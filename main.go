package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/waldhalf/gotmpl/pkg/config"
	"github.com/waldhalf/gotmpl/pkg/driver"
	"github.com/waldhalf/gotmpl/pkg/handlers"
	"github.com/waldhalf/gotmpl/pkg/helpers"
	"github.com/waldhalf/gotmpl/pkg/models"
	"github.com/waldhalf/gotmpl/pkg/render"
	"github.com/waldhalf/gotmpl/pkg/router"
)

var infoLog 	*log.Logger
var errorLog 	*log.Logger

func main() {
	db, err := run()

	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
}

func run() (*driver.DB, error) {
	var app = config.GetConfig()

	// Indicates that we want to put element in Session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// Change to true when on production
	app.InProduction = false

	// Create an info logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	// Create an error logger
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Initializa session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Store session in appConfig
	app.Session = session

	// Connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=postgres")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")


	// Initialize template cache and put it into appconfig var
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}
	app.TemplateCache = tc

	// Initialize repo
	repo := handlers.NewRepo(app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(app)
	helpers.NewHelpers(app)

	fmt.Printf("Starting application on %v\n", app.PortNumber)

	srv := &http.Server{
		Addr: app.PortNumber,
		Handler: router.Routes(app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	return db, nil
}