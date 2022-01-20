package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	print("Hello mod in goLang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	println("Hello mod in golang")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GOlang Series</h1>"))
}
