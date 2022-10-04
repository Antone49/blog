package main

import (
	"errors"
	"log"
	"net/http"
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