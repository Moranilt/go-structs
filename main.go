package main

import (
	"log"
	"net/http"

	"github.com/Moranilt/go-structs/models"
	request "github.com/Moranilt/go-structs/request"
)

func userController(w http.ResponseWriter, r *http.Request) {
	parsedBody := models.User(r).ToJSON()
	w.Write(parsedBody)
}

func locationController(w http.ResponseWriter, r *http.Request) {
	parsedBody := models.Location(r).ToJSON()
	w.Write(parsedBody)
}

func initHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := request.Init(r).CheckMethod(http.MethodPost)
		w.Header().Set("content-type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}
		fn(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", initHandler(userController))
	mux.HandleFunc("/location", initHandler(locationController))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
