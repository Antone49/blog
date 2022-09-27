package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	data "post/data"
)

func (app *Config) addTag(w http.ResponseWriter, r *http.Request) {
	log.Println("addTag")

	var requestPayload data.Tag

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.Insert(requestPayload.Name)
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

	var requestPayload data.Tag

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.PostTag.RemoveByTagId(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.Remove(requestPayload.Id)
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

	var requestPayload data.Tag

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Tag.Update(requestPayload.Id, requestPayload.Name)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updateTag end")
}

func (app *Config) getLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("getLocation")

	var requestPayload data.Location

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	post, err := app.Models.Location.Get(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getLocation end")
}

func (app *Config) addLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("addLocation")

	var requestPayload data.Location

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Location.Insert(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("addLocation end")
}

func (app *Config) removeLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("removeLocation")

	var requestPayload data.Location

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.Models.PostLocation.RemoveByLocationId(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	err = app.Models.Location.Remove(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("removeLocation end")
}

func (app *Config) updateLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("updateLocation")

	var requestPayload data.Location

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Location.Update(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updateLocation end")
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

func (app *Config) addPost(w http.ResponseWriter, r *http.Request) {
	log.Println("addPost")

	var requestPayload data.Post

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Post.Insert(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

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

func (app *Config) updatePostTags(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePostTags")

	var requestPayload struct {
		PostId    int `json:"postId"`
		TagId     []int `json:"tagsId"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.PostTag.RemoveByPostId(requestPayload.PostId)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	err = app.Models.PostTag.InsertFromPostId(requestPayload.PostId, requestPayload.TagId)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updatePostTags end")
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
