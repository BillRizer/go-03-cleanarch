package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	Routes        []Routes
	WebServerPort string
}

type Routes struct {
	Method string
	Path   string
	Handle http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
	s.Routes = append(s.Routes, Routes{
		Method: method,
		Path:   path,
		Handle: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, item := range s.Routes {
		s.Router.Method(item.Method, item.Path, item.Handle)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
