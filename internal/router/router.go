package router

import (
	"github.com/gorilla/mux"
	"github.com/isaro/fish/internal/controller"
	"github.com/isaro/fish/internal/types"
)

var (
	routes = []types.Route{
		{"GET", controller.HealthCheck, "/health", "HealthCheck"},
	}
)

type HttpRouter struct{}

func (r HttpRouter) CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Route).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
