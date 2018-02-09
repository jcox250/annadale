package main

import (
	"log"
	"net/http"
	"os"
	"tech-test-jc/infrastructure"
	"tech-test-jc/interfaces"
	"tech-test-jc/usecases"
)

var DB_TYPE = os.Getenv("DB_TYPE")
var DB_URL = os.Getenv("DATABASE_URL")
var PORT = os.Getenv("PORT")

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbHandler := infrastructure.NewSQLHandler(DB_TYPE, DB_URL)
	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBPersonRepo"] = dbHandler

	personInteractor := usecases.PersonInteractor{}
	personInteractor.PersonRepository = interfaces.NewPersonRepo(handlers)

	personServiceHandler := interfaces.PersonServiceHandler{}
	personServiceHandler.PersonInteractor = personInteractor

	mux := http.NewServeMux()
	http.Handle("/", mux)
	mux.Handle("/people", personServiceHandler)

	srv := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
