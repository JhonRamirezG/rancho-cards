package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jhonrmz/rancho-cards/pkg/config"
	"github.com/jhonrmz/rancho-cards/pkg/handlers"
	"github.com/jhonrmz/rancho-cards/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

//* main is the main func
func main() {

	// Change to true when is production
	app.InProduction = false
	//Create the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	//Store the session int he app configuration session.
	app.Session = session

	tc, err := render.CrateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	//*I store the call to the function CreateTemplateCache in the config file in the type TemplateCache.
	app.TemplateCache = tc
	//* We set this variable to false as a developer mode, which means that we'll read from the disk and not from the cache.
	app.UseCache = false

	//* This is for pass to the handlers the information from config.AppConfig
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	//* This make to our render package access to the configuration file for the app.
	render.NewTemplates(&app)

	//Start the web server.
	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
