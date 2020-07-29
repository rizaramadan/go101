package web

import (
	"net/http"
)

type Route struct {
    Pattern string
    Handler http.HandlerFunc
}

func NewRoute(p string, h http.HandlerFunc) *Route {
    return &Route{p,h}
}

func (r *Route) AddRoute() {
    http.HandleFunc(r.Pattern, r.Handler)
}