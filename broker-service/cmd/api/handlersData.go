package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"io"
	"mime/multipart"
	"bytes"
)

type PointerFuncData func(w http.ResponseWriter, method, command string, r *http.Request)

type ActionDataMap struct {
	VerifyToken     bool
	Function   		PointerFuncData
	Method 			string
	Command         string
}

const MAX_UPLOAD_SIZE = 1024 * 1024 * 5 // 5MB

func (app *Config) HandleDataSubmission(w http.ResponseWriter, r *http.Request) {
	log.Printf("HandleDataSubmission\n")

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Println("MaxBytesReader error : ", err)
		app.errorJSON(w, errors.New("The uploaded file is too big. Please choose an file that's less than 1MB in size"))
		return
	}

	command := r.FormValue("command")
	if len(command) == 0 {
		log.Println("Error : command empty")
		app.errorJSON(w, errors.New("Error : command empty"))
		return
	}

	var actionMap = map[string]ActionDataMap{ 
		"uploadImage" : 		{false, 	app.sendRequestDataImage, "POST", "uploadImage"},
	}

	// if action exists
	if element, ok := actionMap[command]; ok {
		if element.VerifyToken == true {

			token := r.FormValue("token")
			if len(command) == 0 {
				log.Println("Error : command empty")
				app.errorJSON(w, errors.New("Error : command empty"))
				return
			}

			requestPayload := RequestPayload{
				Action: command,
				Token: token,
			}

			err := app.authenticateToken(w, requestPayload)
			if err != nil {
				log.Println("error authenticate token : ", err)
				return
			}

			log.Println("Token verification correct")
		}

		element.Function(w, element.Method, element.Command, r)

	} else {
		log.Println("unknown action : " + command)
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) sendRequestDataImage(w http.ResponseWriter, method, command string, r *http.Request) {
	app.sendRequestAndResponseData(w, method, "http://image-service/" + command, r)
}

func (app *Config) sendRequestAndResponseData(w http.ResponseWriter, method, command string, r *http.Request) {
	log.Println(command)

	err, jsonFromService := app.sendRequestData(w, method, command, r)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	app.sendResponseData(w, command, jsonFromService)

	log.Println(command + " end")
}

func (app *Config) sendRequestData(w http.ResponseWriter, method, command string, r *http.Request) (error, jsonResponse) {
	rdr, hdr, err := r.FormFile("file")
	if err != nil {
		log.Println("Error : ", err)
		return app.errorJSON(w, err) , jsonResponse{}
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", hdr.Filename)
	io.Copy(part, rdr)
	writer.Close()
  
	r, _ = http.NewRequest(method, command, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(r)
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

func (app *Config) sendResponseData(w http.ResponseWriter, command string, jsonFromService jsonResponse) {
	app.writeJSON(w, http.StatusAccepted, jsonFromService)
}