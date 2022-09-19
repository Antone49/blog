package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	data "post/data"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ThumbnailImage string `json:"thumbnailImage",omitempty`
	Image string `json:"image",omitempty`
}

func (app *Config) getAllPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllPosts")
	var requestPayload struct {
		Search    string `json:"search"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	posts, err := data.GetAllPosts(requestPayload.Search)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, posts)

	log.Println("getAllPosts end")
}

func (app *Config) getPost(w http.ResponseWriter, r *http.Request) {
	log.Println("getPost")
	var requestPayload struct {
		Id    int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	post, err := data.GetPost(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getPost end")
}

func (app *Config) getAllTags(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllTags")

	// validate the user against the database
	posts, err := data.GetAllTags()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, posts)

	log.Println("getAllTags end")
}

func (app *Config) getPostTags(w http.ResponseWriter, r *http.Request) {
	log.Println("getPostTags")
	var requestPayload struct {
		Id    int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("id: %d\n", requestPayload.Id)

	// validate the user against the database
	post, err := data.GetPostTags(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getPostTags end")
}

func (app *Config) getLastestPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("getLastestPosts")

	var requestPayload struct {
		Number    int `json:"number"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	post, err := data.GetLastestPosts(requestPayload.Number)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getLastestPosts end")
}

func (app *Config) getAllLocations(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllLocations")

	// validate the user against the database
	locations, err := data.GetAllLocations()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, locations)

	log.Println("getAllLocations end")
}

func (app *Config) sendResponse(w http.ResponseWriter, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New(err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("response good!"),
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
