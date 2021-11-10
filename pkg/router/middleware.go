package router

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/waldhalf/gotmpl/pkg/config"
)
var app = config.GetConfig()

// NoSurf adds csrf protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads ans save session data for current resquest
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}