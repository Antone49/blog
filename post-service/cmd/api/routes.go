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
	mux.Get("/getAllPosts", app.getAllPosts)
	mux.Get("/getPost", app.getPost)
	mux.Post("/addPost", app.addPost)
	mux.Post("/updatePost", app.updatePost)
	mux.Post("/removePost", app.removePost)
	mux.Get("/getLastestPosts", app.getLastestPosts)
	mux.Get("/getPostTags", app.getPostTags)
	mux.Get("/getAllTags", app.getAllTags)
	mux.Get("/getAllLocations", app.getAllLocations)
	mux.Post("/addTag", app.addTag)
	mux.Post("/updateTag", app.updateTag)
	mux.Post("/removeTag", app.removeTag)
	mux.Get("/getLocation", app.getLocation)
	mux.Post("/addLocation", app.addLocation)
	mux.Post("/updateLocation", app.updateLocation)
	mux.Post("/removeLocation", app.removeLocation)
	mux.Post("/updatePostTags", app.updatePostTags)

	return mux
}
