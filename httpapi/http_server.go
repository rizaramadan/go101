package main

import (
	"github.com/labstack/echo"
	""
	"net/http"
)



func main() {
	e := echo.New()
	he


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
