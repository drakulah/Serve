package server

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	pattern regexp.Regexp
	handler func(*http.Request, http.ResponseWriter)
}

type Server struct {
	serverMux *http.ServeMux
	routes    []route
}

func (it Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	path, method := r.URL.Path, strings.ToUpper(r.Method)
	for _, e := range it.routes {
		if e.method == method && e.pattern.MatchString(path) {
			e.handler(r, w)
			return
		}
	}
	w.WriteHeader(400)
	w.Write([]byte("Not Found!"))
}

func (it *Server) Get(pattern string, handler func(*http.Request, http.ResponseWriter)) {
	patternRegExp, err := regexp.Compile(pattern)
	if err != nil {
		newPatternRegExp, err1 := regexp.Compile("^" + pattern + "$")
		if err1 != nil {
			panic(err1)
		}
		patternRegExp = newPatternRegExp
	}
	route := route{
		method:  "GET",
		pattern: *patternRegExp,
		handler: handler,
	}
	it.routes = append(it.routes, route)
}

func (it *Server) Post(pattern string, handler func(*http.Request, http.ResponseWriter)) {
	patternRegExp, err := regexp.Compile(pattern)
	if err != nil {
		newPatternRegExp, err1 := regexp.Compile("^" + pattern + "$")
		if err1 != nil {
			panic(err1)
		}
		patternRegExp = newPatternRegExp
	}
	route := route{
		method:  "POST",
		pattern: *patternRegExp,
		handler: handler,
	}
	it.routes = append(it.routes, route)
}

func (it Server) Listen(addr string, onReady func()) {
	it.serverMux.HandleFunc("/", it.rootHandler)
	onReady()
	if err := http.ListenAndServe(addr, it.serverMux); err != nil {
		log.Fatal(err)
	}
}

func New() Server {
	return Server{
		serverMux: http.NewServeMux(),
	}
}
