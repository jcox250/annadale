package service

import (
	"net/http"

	"github.com/jcox250/annadale/domain"
)

type HomeInteractor interface {
	GetTopNews() ([]domain.Post, error)
}

type HomeService struct {
	Interactor HomeInteractor
}

type HomePageData struct {
	TopNews []domain.Post
}

func NewHomeService(interactor HomeInteractor) *HomeService {
	return &HomeService{
		Interactor: interactor,
	}
}

func (h *HomeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topNews, err := h.Interactor.GetTopNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	data := HomePageData{
		TopNews: topNews,
	}
	templates[homePage].ExecuteTemplate(w, "base", data)
}
