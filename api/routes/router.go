package routes

import (
	"fmt"
	"net/http"
)

type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodPatch   HttpMethod = "PATCH"
)

type Route struct {
	Method  HttpMethod
	Path    string
	Handler http.HandlerFunc
}

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.handleWithMethod("GET", path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.handleWithMethod("POST", path, handler)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.handleWithMethod("DELETE", path, handler)
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.handleWithMethod("PUT", path, handler)
}

func (r *Router) PATCH(path string, handler http.HandlerFunc) {
	r.handleWithMethod("PATCH", path, handler)
}

func (r *Router) HEAD(path string, handler http.HandlerFunc) {
	r.handleWithMethod("HEAD", path, handler)
}

func (r *Router) OPTIONS(path string, handler http.HandlerFunc) {
	r.handleWithMethod("OPTIONS", path, handler)
}

func (r *Router) handleWithMethod(method, path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(fmt.Sprintf("%s %s", method, path), handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
