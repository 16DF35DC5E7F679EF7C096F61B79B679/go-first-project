package main

import (
	"go-first-project/dependencies"
	"net/http"
)


func main() {
	googleSearchCredentials := dependencies.GoogleSearchCredentials{
		ApiKey:         "*********",
		SearchEngineId: "*********",
	}
	app := dependencies.App{
		GoogleSearchCredentials: &googleSearchCredentials,
	}
	router := setupHandlers(&app)
	http.Handle("/", router)
	http.ListenAndServe("localhost:8085", nil)
}