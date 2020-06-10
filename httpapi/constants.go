package main

import "github.com/labstack/echo"

const (
	GET = iota
	POST = iota
	PUT = iota
	DELETE = iota
)

type Endpoint struct {
	Url string
	Method int
	handler func(c echo.Context) error
}

var (
	Endpoints []Endpoint
)