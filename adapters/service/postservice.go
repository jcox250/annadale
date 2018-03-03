package service

import (
	"net/http"
	"strings"

	"github.com/jcox250/annadale/domain"
)

type PostInteractor interface {
	GetPost(id string) domain.Post
	GetAllPosts()
	AddPost(domain.Post) (int64, error)
	GetNews() ([]domain.Post, error)
	GetPages() ([]domain.Post, error)
}

type PostService struct {
	Interactor PostInteractor
}

func NewPostService(interactor PostInteractor) *PostService {
	return &PostService{
		Interactor: interactor,
	}
}

func (p *PostService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.ShowPost(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

func (p *PostService) ShowPost(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	post := p.Interactor.GetPost(id)
	templates[postPage].ExecuteTemplate(w, "base", post)
}
