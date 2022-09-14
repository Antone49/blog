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

	log.Printf("getAllPosts\n")

	// validate the user against the database
	posts, err := data.GetAllPosts()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// var s []Post
	// s = append(s, Post{Id: 1, Title: "Shanghai_VPN", Content: "127.0.0.1azfeeeeeeeeeeeeeeeeeeeeeeeeeeazertyuiopqsdfghjklmwxcvbn123456789azertyuiopqsdfghjklmwxcvbn", Thumbnail: "test.jpeg"})
	// s = append(s, Post{Id: 2, Title: "Beijing_VPN", Content: "127.0.0.2", Thumbnail: "img1.png"})
	
	b, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("json err:", err)
	}
	log.Println(string(b))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", "test"),
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("Post end\n")
}

func (app *Config) getPost(w http.ResponseWriter, r *http.Request) {

	log.Printf("getPost\n")
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
	post, err := data.GetPost(requestPayload.Id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("json err:", err)
	}
	log.Println(string(b))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("getPost!"),
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("getPost end\n")
}

func (app *Config) getPostTags(w http.ResponseWriter, r *http.Request) {

	log.Printf("getPostTags\n")
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

	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("json err:", err)
	}
	log.Println(string(b))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("getPostTags!"),
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("getPostTags end\n")
}

func (app *Config) getLastestPosts(w http.ResponseWriter, r *http.Request) {

	log.Printf("getLastestPosts\n")

	var requestPayload struct {
		Number    int `json:"number"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Printf("r %s\n", r)

	log.Printf("number post to get: %d\n", requestPayload.Number)

	// validate the user against the database
	post, err := data.GetLastestPosts(requestPayload.Number)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println("json err:", err)
	}
	log.Println(string(b))

	payload := jsonResponse{
		Error:   false,
		Message: "getLastestPosts!",
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("getLastestPosts end\n")
}
