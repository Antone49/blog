package main

import (
	"errors"
	"log"
	"net/http"
	data "post/data"
)

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

func (app *Config) addPost(w http.ResponseWriter, r *http.Request) {
	log.Println("addPost")

	// validate the user against the database
	post, err := app.Models.Post.Insert()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("addPost end")
}

func (app *Config) removePost(w http.ResponseWriter, r *http.Request) {
	log.Println("removePost")

	var requestPayload data.Post

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.PostTag.RemoveByPostId(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	err = app.Models.PostLocation.RemoveByPostId(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Post.Delete(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("removePost end")
}

func (app *Config) updatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePost")

	var requestPayload data.Post

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Post.Update(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updatePost end")
}

func (app *Config) updatePostImage(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePostImage")

	var requestPayload data.Post

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Post.UpdateImage(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updatePostImage end")
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