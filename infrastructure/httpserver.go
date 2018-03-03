package infrastructure

import (
	"net/http"

	"github.com/jcox250/annadale/infrastructure/middleware"

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
	logger   middleware.LoggerMiddleware
}

func NewHTTPServer(port string, adapters map[int]http.Handler, logger middleware.LoggerMiddleware) *HTTPServer {
	server := &HTTPServer{
		mux:      http.NewServeMux(),
		port:     port,
		adapters: adapters,
		logger:   logger,
	}
	server.setupRoutes()
	return server
}

func (h *HTTPServer) setupRoutes() {
	h.mux.Handle("/", h.adapters[HomeService])

	// /page/{id}
	h.mux.Handle("/page/", h.adapters[PostService])
	// /admin/addpost
	h.mux.Handle("/admin/", middleware.Authenticate(h.adapters[AdminService]))
	h.mux.Handle("/login/", h.adapters[LoginService])
}

func (h *HTTPServer) Serve() error {
	return http.ListenAndServe(h.port, h.logger.Log(context.ClearHandler(h.mux)))
}
