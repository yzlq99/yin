package yinS

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]http.HandlerFunc),
	}
}

func (e *Engine) addRouter(method string, pattern string, handler http.HandlerFunc) {
	key := method + "_" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler http.HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler http.HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "404 NOT FOUNT: %s\n", req.URL)
	}
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
