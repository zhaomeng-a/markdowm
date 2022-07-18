package router

import (
	"net/http"
	"regexp"
)

type Router struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (router *Router) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.routes[pattern] = handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for pattern, handler := range router.routes {
		if match(pattern, r.URL.Path) {
			//fmt.Println("pattern", pattern, "r", r.URL.Path)
			handler(w, r)
		}
	}
}

func match(pattern, path string) bool {
	tmpPattern := regexp.MustCompile(`:\w+`).ReplaceAllString(pattern, "\\w+")
	matched, _ := regexp.MatchString("^"+tmpPattern+"$", path)
	return matched
}

func New() *Router {
	return &Router{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}
