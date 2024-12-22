package main

import (
	"net/http"

	"com.mal.gotweet.backend/router"
)

func main() {

	routersInit := router.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()
}
