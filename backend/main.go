package main

import (
	"net/http"

	"github.com/CodeSigma/learn-go/router"
)

func main() {

	routersInit := router.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()
}
