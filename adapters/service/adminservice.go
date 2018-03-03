package service

import (
	"net/http"

	"github.com/jcox250/annadale/domain"
)

type AdminInteractor interface {
	GetNews() ([]domain.Post, error)
	GetPages() ([]domain.Post, error)
}

type AdminData struct {
	News  []domain.Post
	Pages []domain.Post
}

type AdminService struct {
	Interactor AdminInteractor
}

func NewAdminService(interactor AdminInteractor) *AdminService {
	return &AdminService{
		Interactor: interactor,
	}
}

func (a *AdminService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	news, err := a.Interactor.GetNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	pages, err := a.Interactor.GetPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	data := AdminData{
		News:  news,
		Pages: pages,
	}

	templates[adminPage].ExecuteTemplate(w, "base", data)
}

func (a *AdminService) HandleLogon(w http.ResponseWriter, r *http.Request) {
	templates[loginPage].ExecuteTemplate(w, "base", nil)
}
