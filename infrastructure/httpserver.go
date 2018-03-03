package infrastructure

import (
	"net/http"
)

const (
	PostService  = "PostService"
	HomeService  = "HomeService"
	AdminService = "AdminService"
)

type HTTPServer struct {
	mux      *http.ServeMux
	port     string
	adapters map[string]http.Handler
}

func NewHTTPServer(port string, adapters map[string]http.Handler) *HTTPServer {
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
}

func (h *HTTPServer) Serve() error {
	return http.ListenAndServe(h.port, h.mux)
}
