package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Post   PostPayload `json:"post,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
		app.getAllPosts(w)
	case "getPost":
		app.getPost(w, requestPayload.Post)
	case "getLastestPosts":
		app.getLastestPosts(w, requestPayload.Post)
	case "getPostTags":
		app.getPostTags(w, requestPayload.Post)
	case "getAllLocations":
		app.getAllLocations(w)
	case "mailContactUs":
		app.mailContactUs(w, requestPayload.Mail)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) mailContactUs(w http.ResponseWriter, mail MailPayload) {
	jsonData, _ := json.MarshalIndent(mail, "", "\t")

	log.Printf("mailContactUs\n")

	request, err := http.NewRequest("POST", "http://mail-service/send", bytes.NewBuffer(jsonData))
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
	payload.Message = "mailContactUs!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("mailContactUs end\n")
}

func (app *Config) getAllPosts(w http.ResponseWriter) {
	log.Printf("getAllPosts\n")

	request, err := http.NewRequest("GET", "http://post-service/getAllPosts", nil)
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
	payload.Message = "getAllPosts!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("getAllPosts end\n")
}

func (app *Config) getPost(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")

	log.Printf("getPost\n")

	request, err := http.NewRequest("GET", "http://post-service/getPost", bytes.NewBuffer(jsonData))
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
	payload.Message = "getPost!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("getPost end\n")
}

func (app *Config) getPostTags(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")

	log.Printf("getPostTags\n")

	request, err := http.NewRequest("GET", "http://post-service/getPostTags", bytes.NewBuffer(jsonData))
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
	payload.Message = "getPostTags!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Println("getPostTags end")
}

func (app *Config) getLastestPosts(w http.ResponseWriter, p PostPayload) {
	jsonData, _ := json.MarshalIndent(p, "", "\t")

	log.Println("getLastestPosts")

	request, err := http.NewRequest("GET", "http://post-service/getLastestPosts", bytes.NewBuffer(jsonData))
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
	payload.Message = "getLastestPosts!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Println("getLastestPosts end")
}

func (app *Config) getAllLocations(w http.ResponseWriter) {
	log.Println("getAllLocations")

	request, err := http.NewRequest("GET", "http://post-service/getAllLocations", nil)
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
	payload.Message = "getAllLocations!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Println("getAllLocations end")
}
