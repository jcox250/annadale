package service

import (
	"fmt"
	"net/http"

	"github.com/jcox250/annadale/domain"
	"github.com/jcox250/annadale/util/session"
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
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	valid, err := a.Interactor.VerifyLogin(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(valid)
	if valid {
		err := session.NewSession(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/", http.StatusFound)
	}
}
