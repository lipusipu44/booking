package main

import (
	"net/http"

	"github.com/anilpatro044/router/pkg/config"
	"github.com/anilpatro044/router/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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