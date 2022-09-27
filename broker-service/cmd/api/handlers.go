package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action 		string  `json:"action"`
	Token 	    string  `json:"token,omitempty"`
	Login       any 	`json:"login,omitempty"`
	Logout      any 	`json:"logout,omitempty"`
	Post   		any 	`json:"post,omitempty"`
	PostTag   	any 	`json:"postTag,omitempty"`
	Tag   		any 	`json:"tag,omitempty"`
	Mail   		any 	`json:"mail,omitempty"`
	Location   	any 	`json:"location,omitempty"`
}

type PointerFunc func(w http.ResponseWriter, method, command string, payload any)

type ActionMap struct {
	VerifyToken     bool
	Function   		PointerFunc
	Method 			string
	Command         string
	Args   			any 
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
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	var actionMap = map[string]ActionMap{ 
		"getAllPosts" : 		{false, 	app.sendRequestPost, 			"GET", 		"getAllPosts", 			requestPayload.Post},
		"getPost" : 			{false, 	app.sendRequestPost, 			"GET", 		"getPost",  			requestPayload.Post},
		"addPost" : 			{true,  	app.sendRequestPost, 			"POST", 	"addPost", 				requestPayload.Post},
		"updatePost" : 			{true,  	app.sendRequestPost,			"POST", 	"updatePost", 			requestPayload.Post},
		"removePost" : 			{true,  	app.sendRequestPost, 			"POST", 	"removePost", 			requestPayload.Post},
		"getLastestPosts" : 	{false, 	app.sendRequestPost, 			"GET", 		"getLastestPosts", 		requestPayload.Post},
		"getPostTags" : 		{false, 	app.sendRequestPost, 			"GET", 		"getPostTags", 			requestPayload.Post},
		"updatePostTags" : 		{true, 		app.sendRequestPost, 			"POST", 	"updatePostTags", 		requestPayload.PostTag},
		"getAllTags" : 			{false, 	app.sendRequestPost, 			"GET", 		"getAllTags", 			nil},
		"getAllLocations" : 	{false, 	app.sendRequestPost, 			"GET", 		"getAllLocations", 		nil},
		"getLocation" : 		{true,  	app.sendRequestPost, 			"GET", 		"getLocation", 			requestPayload.Location},
		"addLocation" : 		{true,  	app.sendRequestPost, 			"POST", 	"addLocation", 			requestPayload.Location},
		"updateLocation" : 		{true,  	app.sendRequestPost,			"POST", 	"updateLocation", 		requestPayload.Location},
		"removeLocation" : 		{true,  	app.sendRequestPost, 			"POST", 	"removeLocation", 		requestPayload.Location},
		"mailContactUs" : 		{false, 	app.sendRequestMail, 			"POST", 	"send", 				requestPayload.Mail},
		"login" : 				{false, 	app.sendRequestAuthentication, 	"POST", 	"login", 				requestPayload.Login},
		"logout" : 				{false, 	app.sendRequestAuthentication, 	"POST", 	"logout", 				requestPayload.Logout},
		"addTag" : 				{true,  	app.sendRequestPost, 			"POST", 	"addTag", 				requestPayload.Tag},
		"updateTag" : 			{true,  	app.sendRequestPost,			"POST", 	"updateTag", 			requestPayload.Tag},
		"removeTag" : 			{true,  	app.sendRequestPost, 			"POST", 	"removeTag", 			requestPayload.Tag},
	}

	// if action exists
	if element, ok := actionMap[requestPayload.Action]; ok {
		if element.VerifyToken == true {
			err := app.authenticateToken(w, requestPayload)
			if err != nil {
				log.Println("error authenticate token : ", err)
				return
			}

			log.Println("Token verification correct")
		}

		element.Function(w, element.Method, element.Command, element.Args)

	} else {
		log.Println("unknown action : " + requestPayload.Action)
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) authenticateToken(w http.ResponseWriter, payload any) error {
	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Println("error json marshall : ", err)
		return err
	}

	err, _ = app.sendRequest(w, "POST", "http://authentication-service/authenticateToken", jsonData)
	
	return err
}

func (app *Config) sendRequestAuthentication(w http.ResponseWriter, method, command string, payload any) {
	app.sendRequestAndResponse(w, method, "http://authentication-service/" + command, payload)
}

func (app *Config) sendRequestPost(w http.ResponseWriter, method, command string, payload any) {
	app.sendRequestAndResponse(w, method, "http://post-service/" + command, payload)
}

func (app *Config) sendRequestMail(w http.ResponseWriter, method, command string, payload any) {
	app.sendRequestAndResponse(w, method, "http://mail-service/" + command, payload)
}

func (app *Config) sendRequestAndResponse(w http.ResponseWriter, method, command string, payload any) {
	log.Println(command)

	var jsonData []byte = nil
	if payload != nil {
		var err error
		jsonData, err = json.MarshalIndent(payload, "", "\t")
		if err != nil {
			log.Println("error json marshall : ", err)
			return
		}
	}

	err, jsonFromService := app.sendRequest(w, method, command, jsonData)
	if err != nil {
		return
	}

	app.sendResponse(w, command, jsonFromService)

	log.Println(command + " end")
}

func (app *Config) sendRequest(w http.ResponseWriter, method, command string, jsonData []byte) (error, jsonResponse) {
	request, err := http.NewRequest(method, command, bytes.NewBuffer(jsonData))
	if err != nil {
		return app.errorJSON(w, err), jsonResponse{}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return app.errorJSON(w, err), jsonResponse{}
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return app.errorJSON(w, errors.New("invalid credentials")), jsonResponse{}
	} else if response.StatusCode != http.StatusAccepted {
		return app.errorJSON(w, errors.New("error call service")), jsonResponse{}
	}

	// create variable we will read response.Body
	var jsonFromService jsonResponse

	// decode
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return app.errorJSON(w, err), jsonResponse{}
	}

	if jsonFromService.Error {
		return app.errorJSON(w, err, http.StatusUnauthorized), jsonResponse{}
	}

	return nil, jsonFromService
}

func (app *Config) sendResponse(w http.ResponseWriter, command string, jsonFromService jsonResponse) {
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}