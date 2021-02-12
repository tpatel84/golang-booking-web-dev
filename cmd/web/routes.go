package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tpatel84/golang-booking-web-dev/pkg/config"
	"github.com/tpatel84/golang-booking-web-dev/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// If we use Pat router (https://github.com/bmizerany/pat)
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// Using Chi router (https://github.com/go-chi/chi)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(writeToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("static", fileServer))

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
