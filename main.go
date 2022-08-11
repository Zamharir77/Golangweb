package main

import (
	"log"
	"net/http"
	"webgolang/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Prosses)

	log.Println("Web starting")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
