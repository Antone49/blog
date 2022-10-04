package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

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
