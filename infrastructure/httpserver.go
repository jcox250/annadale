package infrastructure

import (
	"net/http"

	"github.com/gorilla/context"
)

const (
	PostService = iota
	HomeService
	AdminService
	LoginService
)

type HTTPServer struct {
	mux      *http.ServeMux
	port     string
	adapters map[int]http.Handler
}

func NewHTTPServer(port string, adapters map[int]http.Handler) *HTTPServer {
	server := &HTTPServer{
		mux:      http.NewServeMux(),
		port:     port,
		adapters: adapters,
	}
	server.setupRoutes()
	return server
}

func (h *HTTPServer) setupRoutes() {
	h.mux.Handle("/", h.adapters[HomeService])

	// /page/{id}
	h.mux.Handle("/page/", h.adapters[PostService])
	h.mux.Handle("/admin/", h.adapters[AdminService])
	h.mux.Handle("/login/", h.adapters[LoginService])
}

func (h *HTTPServer) Serve() error {
	return http.ListenAndServe(h.port, context.ClearHandler(h.mux))
}
