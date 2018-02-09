package main

import (
	"log"
	"net/http"
	"tech-test-jc/infrastructure"
	"tech-test-jc/interfaces"
	"tech-test-jc/usecases"
)

const dbType = "mysql"
const connStr = "root:root@(localhost:3306)/techtest"

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbHandler := infrastructure.NewSQLHandler(dbType, connStr)
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
		Addr:    ":5000",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
