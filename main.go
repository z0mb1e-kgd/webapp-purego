package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	message := "Default route"
	w.Write([]byte(message))
}

func snippetCreate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(fmt.Sprintf("Method %s not accepted", req.Method)))
	}
}

func snippetShow(w http.ResponseWriter, req *http.Request) {
	message := "Psst, kid, want some snippet?"
	w.Write([]byte(message))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet", snippetShow)

	serverPort := ":4000"

	log.Println("Server started...")
	defer log.Println("Server shutting down....")
	err := http.ListenAndServe(serverPort, mux)
	log.Fatal(err)
}
