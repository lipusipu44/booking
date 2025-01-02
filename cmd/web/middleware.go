package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func NewCSRFToken(handler http.Handler) http.Handler{
	log.Println("Calling NoCSRF middleware")
	csrfHandler := nosurf.New(handler)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(handler http.Handler) http.Handler{
	log.Println("Load session data")
	return session.LoadAndSave(handler)
}