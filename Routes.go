package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-first-project/dependencies"
	"go-first-project/webRequesthandlers/googleSearch"
	"html/template"
	"net/http"
)

type GoogleSearchHandler struct{
	app *dependencies.App
}

func (googleSearchHandler *GoogleSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	fmt.Println(val)
	fmt.Printf("[GoogleSearchHandlerFunc]Reached here  .. and query is %s \n", val["query"])
	searchQuery := val["query"][0]
	limit := val["limit"][0]
	resultDto, err := googleSearch.Search(searchQuery, limit, googleSearchHandler.app.GoogleSearchCredentials)
	if err != nil {
		fmt.Printf("ResultDTO is "+ resultDto.ToString())
	} else {
		fmt.Printf("Error is : %e ", err)
	}
	fmt.Printf("ResultDTO is "+ resultDto.ToString())
	t, _ := template.ParseFiles("templates/searchResultsPage.gohtml")
	_ = t.Execute(w, resultDto)
}

func HomepageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[HomepageHandlerFunc]Reached here ..")
	t, _ := template.ParseFiles("templates/homepage.gohtml")
	_ = t.Execute(w, "")
	_, _ = fmt.Fprintf(w, "Welcome to the new world")
}

func setupHandlers(app *dependencies.App) *mux.Router{
	r := mux.NewRouter()
	r.Handle("/search/google/", &GoogleSearchHandler{app}).Methods("GET")
	r.HandleFunc("/", HomepageHandlerFunc).Methods("GET")
	return r
}
