package main

import (
	"errors"
	"log"
	"net/http"
)

func (app *Config) getAllPostLocations(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllPostLocations")

	// validate the user against the database
	post, err := app.Models.PostLocation.GetAll()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getAllPostLocations end")
}

func (app *Config) getPostLocations(w http.ResponseWriter, r *http.Request) {
	log.Println("getPostLocations")
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
	post, err := app.Models.PostLocation.GetLocationsFromPostId(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, post)

	log.Println("getPostLocations end")
}

func (app *Config) updatePostLocations(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePostLocations")

	var requestPayload struct {
		PostId    int `json:"postId"`
		LocationId     []int `json:"locationId"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.PostLocation.RemoveByPostId(requestPayload.PostId)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	err = app.Models.PostLocation.InsertFromPostId(requestPayload.PostId, requestPayload.LocationId)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updatePostLocations end")
}