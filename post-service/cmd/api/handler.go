package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ThumbnailImage string `json:"thumbnailImage",omitempty`
	Image string `json:"image",omitempty`
}

func (app *Config) addTag(w http.ResponseWriter, r *http.Request) {
	log.Println("addTag")

	var requestPayload struct {
		Id    string `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.Insert(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("addTag end")
}

func (app *Config) removeTag(w http.ResponseWriter, r *http.Request) {
	log.Println("removeTag")

	var requestPayload struct {
		Id    string `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.PostTag.RemoveByTagName(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.RemoveByName(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("removeTag end")
}

func (app *Config) updateTag(w http.ResponseWriter, r *http.Request) {
	log.Println("updateTag")

	var requestPayload struct {
		OldName    string `json:"oldName"`
		NewName    string `json:"newName"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.UpdateByName(requestPayload.OldName, requestPayload.NewName)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updateTag end")
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
	posts, err := app.Models.Post.GetAll(requestPayload.Search)
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
	post, err := app.Models.Post.Get(requestPayload.Id)
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
	posts, err := app.Models.Tag.GetAll()
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
	post, err := app.Models.PostTag.GetTagsFromPostId(requestPayload.Id)
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
	post, err := app.Models.Post.GetLastest(requestPayload.Number)
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
	locations, err := app.Models.Location.GetAll()
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
