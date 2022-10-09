package main

import (
	"errors"
	"log"
	"net/http"
	data "message/data"
)

func (app *Config) getAll(w http.ResponseWriter, r *http.Request) {
	log.Println("getAll")

	// validate the user against the database
	messages, err := app.Models.Message.GetAll()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, messages)

	log.Println("getAll end")
}

func (app *Config) getAllFromPostId(w http.ResponseWriter, r *http.Request) {
	log.Println("getAllFromPostId")

	var requestPayload data.Message

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	message, err := app.Models.Message.GetAllFromPostId(requestPayload.PostId)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, message)

	log.Println("getAllFromPostId end")
}

func (app *Config) add(w http.ResponseWriter, r *http.Request) {
	log.Println("add")

	var requestPayload data.Message

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Message.Insert(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	resp := jsonResponse{
		Error: false,
	}

	app.sendResponse(w, resp)

	log.Println("add end")
}

func (app *Config) remove(w http.ResponseWriter, r *http.Request) {
	log.Println("remove")

	var requestPayload data.Message

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Message.Delete(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("remove end")
}

func (app *Config) update(w http.ResponseWriter, r *http.Request) {
	log.Println("update")

	var requestPayload data.Message

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Message.Update(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("update end")
}

func (app *Config) updateValidity(w http.ResponseWriter, r *http.Request) {
	log.Println("updateValidity")

	var requestPayload data.Message

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	err = app.Models.Message.UpdateValidity(requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, nil)

	log.Println("updateValidity end")
}

func (app *Config) getLastest(w http.ResponseWriter, r *http.Request) {
	log.Println("getLastest")

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
	message, err := app.Models.Message.GetLastest(requestPayload.Number)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	app.sendResponse(w, message)

	log.Println("getLastest end")
}