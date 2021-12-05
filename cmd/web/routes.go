package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jhonrmz/rancho-cards/pkg/config"
	"github.com/jhonrmz/rancho-cards/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	//* Call to the midleware.
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//* This is where we set the routes from out application using the package chi.
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/shop", handlers.Repo.Shop)
	mux.Get("/offers", handlers.Repo.Offers)
	mux.Get("/orders", handlers.Repo.Orders)

	return mux
}
