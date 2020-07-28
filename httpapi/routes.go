package httpapi

import (
    "net/http"
)

type Route struct {
    Pattern string
    Handler HandlerFunc
}

func (r *Route) Add() {
    http.HandlerFunc(r.Pattern, r.Handler);
}
