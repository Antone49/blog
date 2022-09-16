package main

import (
	"log"
	"net/http"
)


func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	log.Println("SendMail")

	type mailMessage struct {
		Email 	string `json:"email"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJSON(w, r, & requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	// insert data
	msg := Message{
		From: requestPayload.Email,
		To: "test@ded.fr",
		Subject: requestPayload.Subject,
		Data: requestPayload.Message,
	}

	err = app.Mailer.SendSMPTMessage(msg)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "Email " + requestPayload.Email,
	}

	app.writeJSON(w, http.StatusAccepted, resp)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	log.Println("SendMail end")
}