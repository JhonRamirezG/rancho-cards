package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//* NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// * SessionLOad loads and save session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
