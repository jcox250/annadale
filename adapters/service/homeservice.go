package service

import (
	"net/http"
)

type HomeInteractor interface {
}

type HomeService struct {
	Interactor HomeInteractor
}

func NewHomeService(interactor HomeInteractor) *HomeService {
	return &HomeService{
		Interactor: interactor,
	}
}

func (h *HomeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates[homePage].ExecuteTemplate(w, "base", nil)
}
