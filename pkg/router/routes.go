package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/waldhalf/gotmpl/pkg/config"
	"github.com/waldhalf/gotmpl/pkg/handlers"
)


func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Initialize middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Web routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
