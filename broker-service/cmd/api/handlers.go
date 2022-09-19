package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action 		string      	`json:"action"`
	Auth   		AuthPayload 	`json:"auth,omitempty"`
	AllPosts   	AllPostsPayload `json:"allPosts,omitempty"`
	Post   		PostPayload 	`json:"post,omitempty"`
	Mail   		MailPayload 	`json:"mail,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AllPostsPayload struct {
	Search string `json:"search,omitempty"`
}

type PostPayload struct {
	Id     int `json:"id,omitempty"`
	Number int `json:"number,omitempty"`
}

type MailPayload struct {
	Email 	string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

	log.Printf("Broker\n")
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {

	log.Printf("HandleSubmission\n")

	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "getAllPosts":
		app.getAllPosts(w, requestPayload.AllPosts)
	case "getPost":
		app.getPost(w, requestPayload.Post)
	case "getLastestPosts":
		app.getLastestPosts(w, requestPayload.Post)
	case "getPostTags":
		app.getPostTags(w, requestPayload.Post)
	case "getAllTags":
		app.getAllTags(w)
	case "getAllLocations":
		app.getAllLocations(w)
	case "mailContactUs":
		app.mailContactUs(w, requestPayload.Mail)
	default:
		log.Println("unknown action : " + requestPayload.Action)
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) mailContactUs(w http.ResponseWriter, mail MailPayload) {
	jsonData, _ := json.MarshalIndent(mail, "", "\t")
	app.sendRequestPost(w, "POST", "send", jsonData)
}

func (app *Config) getAllPosts(w http.ResponseWriter, allPosts AllPostsPayload) {
	jsonData, _ := json.MarshalIndent(allPosts, "", "\t")
	app.sendRequestPost(w, "GET", "getAllPosts", jsonData)
}

func (app *Config) getPost(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	app.sendRequestPost(w, "GET", "getPost", jsonData)
}

func (app *Config) getPostTags(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	app.sendRequestPost(w, "GET", "getPostTags", jsonData)
}

func (app *Config) getAllTags(w http.ResponseWriter) {
	app.sendRequestPost(w, "GET", "getAllTags", nil)
}

func (app *Config) getLastestPosts(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")
	app.sendRequestPost(w, "GET", "getLastestPosts", jsonData)
}

func (app *Config) getAllLocations(w http.ResponseWriter) {
	app.sendRequestPost(w, "GET", "getAllLocations", nil)
}

func (app *Config) sendRequestPost(w http.ResponseWriter, method, command string, jsonData []byte) {
	app.sendRequest(w, method, "http://post-service/" + command, jsonData)
}

func (app *Config) sendRequestMail(w http.ResponseWriter, method, command string, jsonData []byte) {
	app.sendRequest(w, method, "http://mail-service/" + command, jsonData)
}

func (app *Config) sendRequest(w http.ResponseWriter, method, command string, jsonData []byte) {
	log.Println(command)

	request, err := http.NewRequest(method, command, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	var responsePayload jsonResponse
	if err != nil {
		log.Println(responsePayload.Message)
	}

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error call auth service"))
		return
	}

	// create variable we will read response.Body
	var jsonFromService jsonResponse

	// decode
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = command
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Println(command + " end")
}