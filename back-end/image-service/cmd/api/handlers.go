package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"time"
	"io"
	"encoding/json"
	"errors"
)

type jsonImage struct {
	Filename string `json:"filename"`
}

const DATA_DIRECTORY = "usr/src"
const MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10MB

func (app *Config) sendResponse(w http.ResponseWriter, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New(err.Error()), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("response good!"),
		Data: string(b),
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UploadImage(w http.ResponseWriter, r *http.Request) {
	log.Println("uploadImage")

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Println("MaxBytesReader error : ", err)
		app.errorJSON(w, err)
		return
	}

		// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println("FormFile error : ", err)
		app.errorJSON(w, err)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll(DATA_DIRECTORY, os.ModePerm)
	if err != nil {
		log.Println("MkdirAll error : ", err)
		app.errorJSON(w, err)
		return
	}

	filepath := fmt.Sprintf("/images/%d_%s", time.Now().Unix(), fileHeader.Filename)

	// Create a new file in the uploads directory
	dst, err := os.Create(DATA_DIRECTORY + filepath)
	if err != nil {
		log.Println("Create error : ", err)
		app.errorJSON(w, err)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println("Copy error : ", err)
		app.errorJSON(w, err)
		return
	}

	payload := jsonImage{
		Filename:   filepath,
	}

	app.sendResponse(w, payload)

	log.Println("uploadImage end")
}