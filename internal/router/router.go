package router

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request)

type Router interface {
	Get(path string, handler Handler)
	Post(path string, handler Handler)
	Put(path string, handler Handler)
	Delete(path string, handler Handler)

	Group(group string) Router

	ServeHTTP(http.ResponseWriter, *http.Request)
}

type router struct {
	*mux.Router
}

func NewRouter() Router {
	return &router{
		mux.NewRouter(),
	}
}

func (r *router) Get(path string, handler Handler) {
	r.Methods("GET").HandlerFunc(handler)
}

func (r *router) Post(path string, handler Handler) {
	r.Methods("POST").HandlerFunc(handler)
}

func (r *router) Delete(path string, handler Handler) {
	r.Methods("DELETE").HandlerFunc(handler)
}

func (r *router) Put(path string, handler Handler) {
	r.Methods("PUT").HandlerFunc(handler)
}

func (r *router) Group(group string) Router {
	return &router{
		r.NewRoute().PathPrefix("/" + strings.Trim(group, "/")).Subrouter(),
	}
}
