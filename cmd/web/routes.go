package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/lipusipu44/booking/pkg/config"
	"github.com/lipusipu44/booking/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{
	mux:=chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NewCSRFToken)
	mux.Use(SessionLoad)
	mux.Get("/home",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)

	return mux
}