package main

import (
	"log"
	"net/http"
	"tech-test-jc/infrastructure"
	"tech-test-jc/interfaces"
	"tech-test-jc/usecases"
)

const dbType = "mysql"
const connStr = "root:root@(localhost:3306)/flexbdna"

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbHandler := infrastructure.NewSQLHandler(dbType, connStr)
	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBPersonRepo"] = dbHandler

	pi := usecases.PersonInteractor{}
	pi.PersonRepository = interfaces.NewPersonRepo(handlers)

	psh := interfaces.PersonServiceHandler{}
	psh.PersonInteractor = pi

	mux := http.NewServeMux()
	http.Handle("/", mux)
	mux.HandleFunc("/people", psh.ShowPeople)

	srv := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
