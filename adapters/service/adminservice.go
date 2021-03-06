package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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
	case http.MethodPut:
		a.updatePost(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (a *AdminService) handleRouting(w http.ResponseWriter, r *http.Request) {
	route := strings.Split(r.URL.Path, "/")[2]

	switch route {
	case "":
		a.showAdminPage(w, r)
		return
	case "editpost":
		a.showAddPostPage(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		if err := templates[notFoundPage].ExecuteTemplate(w, "content", nil); err != nil {
			http.Redirect(w, r, "/500", http.StatusInternalServerError)
			return
		}
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
	if err := templates[adminPage].ExecuteTemplate(w, "base", data); err != nil {
		http.Redirect(w, r, "/500", http.StatusInternalServerError)
	}
}

func (a *AdminService) showAddPostPage(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	data := domain.Post{}

	if len(path) == 4 {
		id := path[3]
		data = a.Interactor.GetPost(id)
	}
	templates[editPostPage].ExecuteTemplate(w, "base", data)
}

func (a *AdminService) addPost(w http.ResponseWriter, r *http.Request) {
	rowsAdded, err := a.Interactor.AddPost(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rowsAdded > 0 {
		// Will update this to redirect to a 'successfuly updated page'
		http.Redirect(w, r, "/admin/", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/genericErrorMessage", http.StatusFound)
}

func (a *AdminService) updatePost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	updated, err := a.Interactor.EditPost(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if updated {
		// Will update this to redirect to a 'successfuly updated page'
		http.Redirect(w, r, "/admin/", http.StatusFound)
	}
	http.Redirect(w, r, "/genericErrorMessage", http.StatusFound)
}
