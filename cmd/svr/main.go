package main

import (
	"log"
	"net/http"
	"os"
)

var config = "config.json"

func main() {
	//load the config here
	appConfig := config.Load
	//handler, err := establishFacade(config, gitCommit)

	if err != nil {
		log.Printf("failed to set up the service on the service", err)
		os.Exit(1)
	}

	//establish router mux

	err := http.ListenAndServe("3000", mux)
}

func index(r *http.Request, w http.ResponseWriter) {

}
