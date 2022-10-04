package main

import (
	"errors"
	"log"
	"net/http"
	"post/data"
)

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
