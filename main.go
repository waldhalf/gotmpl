package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/waldhalf/gotmpl/pkg/config"
	"github.com/waldhalf/gotmpl/pkg/handlers"
	"github.com/waldhalf/gotmpl/pkg/render"
	"github.com/waldhalf/gotmpl/pkg/router"
)


func main() {
	var app = config.GetConfig()

	// Change to true when on production
	app.InProduction = false

	// Initializa seesion
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Store session in appConfig
	app.Session = session

	// Initialize template cache and put it into appconfig var
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	// Initialize repo
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	render.NewTemplates(app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application on %v\n", app.PortNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: app.PortNumber,
		Handler: router.Routes(app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}