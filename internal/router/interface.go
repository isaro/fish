package router

import "github.com/gorilla/mux"

type Router interface {
	CreateRouter() *mux.Router
}
