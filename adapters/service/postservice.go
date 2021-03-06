package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/jcox250/annadale/domain"
)

type PostInteractor interface {
	GetPost(id string) domain.Post
	GetAllPosts()
	AddPost(r *http.Request) (int64, error)
	EditPost(r *http.Request) (bool, error)
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
	if err := templates[postPage].ExecuteTemplate(w, "base", post); err != nil {
		log.Println(err)
		http.Redirect(w, r, "/500", http.StatusInternalServerError)
	}
}
