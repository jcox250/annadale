package personcontroller

import (
	"bike-api/util/httputil"
	"encoding/json"
	"flexbdna/model/person"
	"log"
	"net/http"
)

type PersonController person.Person

func (p PersonController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAll(w, r)
		break
	case http.MethodPost:
		post(w, r)
		break
	default:
		httputil.NotImplemented(w)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	people, err := person.GetAll()
	if err != nil {
		httputil.InternalServerError(w, err)
	}

	data, err := json.Marshal(people)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}
	httputil.WriteResponse(w, data)
}

func post(w http.ResponseWriter, r *http.Request) {
	var p person.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("Can't decode json into person: %v", err)
	}

	rowsAffected, err := person.Create(p)
	if err != nil {
		httputil.InternalServerError(w, err)
	}

	added := rowsAffected == 1

	if added {
		w.WriteHeader(http.StatusOK)
	} else {
		httputil.InternalServerError(w, err)
	}

}
