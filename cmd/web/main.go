package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tpatel84/golang-booking-web-dev/internal/config"
	"github.com/tpatel84/golang-booking-web-dev/internal/handlers"
	"github.com/tpatel84/golang-booking-web-dev/internal/helpers"
	"github.com/tpatel84/golang-booking-web-dev/internal/models"
	"github.com/tpatel84/golang-booking-web-dev/internal/renders"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT = ":10000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger


func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Starting web application on port %s", PORT))

	//_ = http.ListenAndServe(PORT, nil)

	srv := &http.Server{
		Addr: PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start an application : ", err)
	}
}

func run() error {
	// Store data in session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	//Setup an logger for application
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create a template cache")
		return err
	}
	app.TemplateCache = tc

	// Set the UseCache to false
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	helpers.NewHelpers(&app)

	return nil
}