package main

import (
	"log"
	"net/http"
	"tech-test-jc/interfaces"
)

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mux := http.NewServeMux()
	http.Handle("/", mux)

	psh := interfaces.PersonServiceHandler{}
	mux.HandleFunc("/people", psh.ShowPeople)

	srv := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
