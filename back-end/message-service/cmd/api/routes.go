package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/getAll", app.getAll)
	mux.Get("/getAllFromPostId", app.getAllFromPostId)
	mux.Post("/add", app.add)
	mux.Post("/update", app.update)
	mux.Post("/updateValidity", app.updateValidity)
	mux.Post("/remove", app.remove)
	mux.Get("/getLastest", app.getLastest)

	return mux
}
