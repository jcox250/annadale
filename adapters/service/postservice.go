package service

import (
	"fmt"
	"net/http"
	"strings"
)

type PostInteractor interface {
	GetPost(id string)
	GetAllPosts()
}

type PostService struct {
	Interactor PostInteractor
}

func NewPostAdapter(interactor PostInteractor) *PostService {
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
	id := strings.Split(r.URL.Path, "/")[0]
	fmt.Println(id)
	// p.Interactor.GetPost(id)
}
