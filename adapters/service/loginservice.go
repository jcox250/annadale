package service

import (
	"fmt"
	"log"
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
	switch r.Method {
	case http.MethodGet:
		sessionExists, err := session.SessionExists(r, session.SessionName)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		if sessionExists {
			http.Redirect(w, r, "/admin/", http.StatusFound)
			return
		}
		a.showLoginPage(w, r)
		return
	case http.MethodPost:
		a.handleLogin(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (a *LoginService) showLoginPage(w http.ResponseWriter, r *http.Request) {
	templates[loginPage].ExecuteTemplate(w, "base", nil)
}

func (a *LoginService) handleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := domain.User{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	valid, err := a.Interactor.VerifyLogin(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(valid)
	if valid {
		err := session.NewSession(w, r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/", http.StatusFound)
	}
}
