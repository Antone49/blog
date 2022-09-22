package main

import (
	"bytes"
	"encoding/json"
	"time"
	"errors"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var tokenDuration = time.Duration(6*time.Hour)

func (app *Config) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login\n")

	var requestPayload struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByName(requestPayload.Name)
	if err != nil {
		log.Println("error GetByName :", err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid password"), http.StatusBadRequest)
		return
	}

	randomToken := make([]byte, 32)
    _, err = rand.Read(randomToken)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("invalid token generation"), http.StatusBadRequest)
		return     
	}  

    authToken := base64.URLEncoding.EncodeToString(randomToken)

	// insert the token in the database
	_, err = app.Models.Token.Insert(authToken, user.ID, time.Duration(tokenDuration))
	if err != nil {
		log.Println("error: ", err)
		app.errorJSON(w, errors.New("invalid token insertion"), http.StatusBadRequest)
		return
	}

	// log authentication
	err = app.logRequest("login", fmt.Sprintf("%s logged in", user.Name))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Name),
		Data:    user,
		Token:   authToken,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("Login end\n")
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Logout\n")

	var requestPayload struct {
		Token   	string `json:"token"`
		Username 	string `json:"user"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	pattern := "Bearer "
	token := requestPayload.Token
	if strings.Contains(requestPayload.Token, pattern) {
		token = requestPayload.Token[len(pattern):len(requestPayload.Token)]
	}

	// Supprime Token
	app.Models.Token.DeleteByToken(token)
	if err != nil {
		log.Println("error DeleteByToken :", err)
	}

	user, err := app.Models.User.GetByName(requestPayload.Username)
	if err != nil {
		log.Println("error GetByName :", err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// Supprime tous les tokens presents d'un user
	app.Models.Token.DeleteByUserId(user.ID)

	// log logout
	err = app.logRequest("logout", fmt.Sprintf("%s logout in", user.Name))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged out user %s", user.Name),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("Logout end\n")
}

func (app *Config) AuthenticateToken(w http.ResponseWriter, r *http.Request) {
	log.Printf("AuthenticateToken\n")

	var requestPayload struct {
		Token     string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	pattern := "Bearer "
	token := requestPayload.Token
	if strings.Contains(requestPayload.Token, pattern) {
		token = requestPayload.Token[len(pattern):len(requestPayload.Token)]
	}

	// validate the user against the database
	authToken, err := app.Models.Token.GetByToken(token)
	if err != nil {
		log.Println("error GetByToken :", err)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	if time.Now().After(authToken.ExpiredAt) {
		authToken.Delete()

		err := "Authorization token expired"
		log.Println("error : ", err, " remove token (", token, ") for userId ", authToken.UserId)
		app.errorJSON(w, errors.New(err), http.StatusUnauthorized)
		return
	}

	log.Println("Authorization token correct for the userId ", authToken.UserId)

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("AuthenticateToken!"),
	}

	app.writeJSON(w, http.StatusAccepted, payload)

	log.Printf("AuthenticateToken end\n")
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceUrl := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)

	return nil
}
