package interfaces

import (
	"encoding/json"
	"net/http"
	"tech-test-jc/domain"
	"tech-test-jc/util/httputil"
)

// PersonInteractor does y
type PersonInteractor interface {
	People([]domain.Person, error)
}

// PersonServiceHandler foo
type PersonServiceHandler struct {
	PersonInteractor PersonInteractor
}

// ShowPeople will do y
func (p PersonServiceHandler) ShowPeople(w http.ResponseWriter, r *http.Request) {
	people := []domain.Person{
		domain.Person{
			Name: "James",
		},
		domain.Person{
			Name: "Alan",
		},
		domain.Person{
			Name: "Nora",
		},
	}

	b, err := json.Marshal(people)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	httputil.WriteResponse(w, b)
}
