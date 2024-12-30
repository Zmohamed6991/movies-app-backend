package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)
	mux.Get("/movies", app.allMovies)


	return mux


}