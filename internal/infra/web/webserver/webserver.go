package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	if method == "GET" {
		s.Router.Get(path, handler)
	}
	if method == "POST" {
		s.Router.Post(path, handler)

	}
	//s.Router.HandleFunc(path, handler)
}

func (s *WebServer) RegisterMiddleware() {
	s.Router.Use(middleware.Logger)
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {

	http.ListenAndServe(s.WebServerPort, s.Router)
}
