package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitinIKEA/chatbot-backend-rest/internal/config"
	"github.com/nitinIKEA/chatbot-backend-rest/internal/db"
)

type Service struct {
	Router  *mux.Router
	DBConns db.DBConns
	Conf    *config.Conf
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (s *Service) NewRouter() {
	var routes = s.PrepareRoutes()
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	s.Router = router
}

func (s *Service) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
