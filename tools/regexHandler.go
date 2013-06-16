package tools

import (
	"net/http"
	"regexp"
)

// inspired by the following sources with some small changes:
//http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc
//https://github.com/raymi/quickerreference

type route struct {
    pattern *regexp.Regexp
    handler http.Handler
}

type RegexpHandler struct {
    routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
    h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(strPattern string, handler func(http.ResponseWriter, *http.Request)) {
	// encapsulate string pattern with start and end constraints
	// so that HandleFunc would work as for Python GAE
	pattern := regexp.MustCompile("^"+strPattern+"$")
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
        if route.pattern.MatchString(r.URL.Path) {
            route.handler.ServeHTTP(w, r)
            return
        }
    }
    // no pattern matched; send 404 response
    http.NotFound(w, r)
}
