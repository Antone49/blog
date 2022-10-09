package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort  = "80"
)

type Config struct {
}

func main() {

	app := Config{
	}

	log.Println("Sarting image service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}