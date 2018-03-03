package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
	Interactor PostInteractor
}

func NewAdminService(interactor PostInteractor) *AdminService {
	return &AdminService{
		Interactor: interactor,
	}
}

func (a *AdminService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.handleRouting(w, r)
		return
	case http.MethodPost:
		a.addPost(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (a *AdminService) handleRouting(w http.ResponseWriter, r *http.Request) {
	route := strings.Split(r.URL.Path, "/")[2]

	switch route {
	case "":
		fmt.Println("empty route")
		a.showAdminPage(w, r)
		return
	case "addpost":
		fmt.Println("addpost route")
		a.showAddPostPage(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		templates[notFoundPage].ExecuteTemplate(w, "content", nil)
		return
	}

}

func (a *AdminService) showAdminPage(w http.ResponseWriter, r *http.Request) {
	news, err := a.Interactor.GetNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pages, err := a.Interactor.GetPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := AdminData{
		News:  news,
		Pages: pages,
	}
	templates[adminPage].ExecuteTemplate(w, "base", data)
}

func (a *AdminService) showAddPostPage(w http.ResponseWriter, r *http.Request) {
	templates[addPostPage].ExecuteTemplate(w, "base", nil)
}

func (a *AdminService) addPost(w http.ResponseWriter, r *http.Request) {
	post := domain.Post{
		ID:          "3", //will be uuid
		Title:       r.PostFormValue("title"),
		Content:     r.PostFormValue("content"),
		AddedBy:     "1",
		AddedDate:   time.Now(),
		UpdatedBy:   "1",
		UpdatedDate: time.Now(),
		Archive:     false,
	}

	rowsAdded, err := a.Interactor.AddPost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rowsAdded > 0 {
		w.WriteHeader(http.StatusOK)
	}
}
