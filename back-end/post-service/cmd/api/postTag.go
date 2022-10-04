package main

import (
	"errors"
	"log"
	"net/http"
)

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
