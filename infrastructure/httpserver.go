package infrastructure

import (
	"net/http"
)

const (
	PostService = "PostService"
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
	h.mux.Handle("/page/", h.adapters[PostService])
}

func (h *HTTPServer) Serve() error {
	return http.ListenAndServe(h.port, h.mux)
}
