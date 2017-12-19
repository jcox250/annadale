package main

import (
	"flexbdna/controller/personcontroller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var pc personcontroller.PersonController

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/people", pc)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
