package main

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/rizaramadan/"
)

type Method int

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

func main() {
	e := echo.New()

	for _, elm := range Endpoints {
		switch elm.Method {
			case GET:
				e.GET(elm.Url, elm.handler)
			case POST:
				e.POST(elm.Url, elm.handler)
			case PUT:
				e.PUT(elm.Url, elm.handler)
			case DELETE:
				e.DELETE(elm.Url, elm.handler)
			default:
				panic("error")
		}
	}


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))


}

func saveUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
