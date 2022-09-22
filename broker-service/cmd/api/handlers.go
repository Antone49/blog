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
	Tag   		any 	`json:"tag,omitempty"`
	Mail   		any 	`json:"mail,omitempty"`
}

type PointerFunc func(w http.ResponseWriter, payload any)

type ActionMap struct {
	VerifyToken     bool
	Function   		PointerFunc 
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
		"getAllPosts" : 		{false, 	app.getAllPosts, 		requestPayload.Post},
		"getPost" : 			{false, 	app.getPost,  			requestPayload.Post},
		"getLastestPosts" : 	{false, 	app.getLastestPosts, 	requestPayload.Post},
		"getPostTags" : 		{false, 	app.getPostTags, 		requestPayload.Post},
		"getAllTags" : 			{false, 	app.getAllTags, 		nil},
		"getAllLocations" : 	{false, 	app.getAllLocations, 	nil},
		"mailContactUs" : 		{false, 	app.mailContactUs, 		requestPayload.Mail},
		"login" : 				{false, 	app.login, 				requestPayload.Login},
		"logout" : 				{false, 	app.logout, 			requestPayload.Logout},
		"addTag" : 				{true,  	app.addTag, 			requestPayload.Tag},
		"updateTag" : 			{true,  	app.updateTag, 			requestPayload.Tag},
		"removeTag" : 			{true,  	app.removeTag, 			requestPayload.Tag},
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

		element.Function(w, element.Args)

	} else {
		log.Println("unknown action : " + requestPayload.Action)
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) removeTag(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "POST", "removeTag", payload)
}

func (app *Config) updateTag(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "POST", "updateTag", payload)
}

func (app *Config) addTag(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "POST", "addTag", payload)
}

func (app *Config) login(w http.ResponseWriter, payload any) {
	app.sendRequestAuthentication(w, "POST", "login", payload)
}

func (app *Config) logout(w http.ResponseWriter, payload any) {
	app.sendRequestAuthentication(w, "POST", "logout", payload)
}

func (app *Config) mailContactUs(w http.ResponseWriter, payload any) {
	app.sendRequestMail(w, "POST", "send", payload)
}

func (app *Config) getAllPosts(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getAllPosts", payload)
}

func (app *Config) getPost(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getPost", payload)
}

func (app *Config) getPostTags(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getPostTags", payload)
}

func (app *Config) getAllTags(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getAllTags", payload)
}

func (app *Config) getLastestPosts(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getLastestPosts", payload)
}

func (app *Config) getAllLocations(w http.ResponseWriter, payload any) {
	app.sendRequestPost(w, "GET", "getAllLocations", payload)
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