package main

import (
	"log"
	"net/http"

	request "github.com/Moranilt/go-structs/request"
)

func initFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	parseBody := request.ParsedBody(r)
	w.Write(request.ToJSON(parseBody))
}

func main() {
	http.HandleFunc("/", initFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
