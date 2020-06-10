package hello

import (
	"github.com/labstack/echo"
	"net/http"
)

func Init(e *echo.Echo) {
	e.GET("/", HelloWorld)
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}




