package interfaces

import (
	"encoding/json"
	"net/http"
	"tech-test-jc/domain"
	"tech-test-jc/util/httputil"
)

type PersonInteractor interface {
	People() ([]domain.Person, error)
	AddPerson(domain.Person) (int64, error)
}

type PersonServiceHandler struct {
	PersonInteractor PersonInteractor
}

func (p PersonServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.ShowPeople(w, r)
		break
	case http.MethodPost:
		p.AddPerson(w, r)
		break
	default:
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

func (p PersonServiceHandler) ShowPeople(w http.ResponseWriter, r *http.Request) {
	people, err := p.PersonInteractor.People()
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	b, err := json.Marshal(people)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	httputil.WriteResponse(w, b)
}

func (p PersonServiceHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	var person domain.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	rowsUpdated, err := p.PersonInteractor.AddPerson(person)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	if rowsUpdated == 1 {
		b := []byte("record added")
		w.Write(b)
	}
}
