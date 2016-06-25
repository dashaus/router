package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/monoculum/revr"
)

type Route struct {
	path   string
	router *Router
}

type Router struct {
	*httprouter.Router
	store *revr.URLStore
}

// Force New() to implement the http.Handler interface
var _ http.Handler = New()

func New() *Router {
	m := httprouter.New()
	s := revr.New()
	return &Router{
		Router: m,
		store:  s,
	}
}

func (r *Router) GET(path string, handle httprouter.Handle) *Route {
	r.Handle("GET", path, handle)
	route := &Route{path, r}
	return route
}

func (r *Router) PUT(path string, handle httprouter.Handle) *Route {
	r.Handle("PUT", path, handle)
	route := &Route{path, r}
	return route
}

func (r *Router) POST(path string, handle httprouter.Handle) *Route {
	r.Handle("POST", path, handle)
	route := &Route{path, r}
	return route
}

func (r *Router) PATCH(path string, handle httprouter.Handle) *Route {
	r.Handle("PATCH", path, handle)
	route := &Route{path, r}
	return route
}

func (r *Router) DELETE(path string, handle httprouter.Handle) *Route {
	r.Handle("DELETE", path, handle)
	route := &Route{path, r}
	return route
}

func (r *Router) ServeFiles(path string, root http.FileSystem) {
	r.ServeFiles(path, root)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}

func (r *Router) MustReverse(name string, params ...string) string {
	return r.store.MustReverse(name, params...)
}

func (r *Route) Name(name string) *Route {
	r.router.store.MustAdd(name, r.path)
	return r
}
