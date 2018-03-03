package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jcox250/annadale/domain"
)

type LoginInteractor interface {
	VerifyLogin(domain.User) (bool, error)
}

type LoginService struct {
	Interactor LoginInteractor
}

func NewLoginService(interactor LoginInteractor) *LoginService {
	return &LoginService{
		Interactor: interactor,
	}
}

func (a *LoginService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		templates[loginPage].ExecuteTemplate(w, "base", nil)
		return
	case http.MethodPost:
		a.handleLogin(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (a *LoginService) handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := domain.User{
		Username: strings.Join(r.Form["username"], ""),
		Password: strings.Join(r.Form["password"], ""),
	}

	valid, err := a.Interactor.VerifyLogin(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(valid)
	if valid {
		http.Redirect(w, r, "/admin/", http.StatusFound)
		// set session
		// redirect to admin
	}
}
