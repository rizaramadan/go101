package main

import (
	"github.com/labstack/echo"
	"github.com/rizaramadan/go101/httpapi/hello"
	"net/http"
)



func main() {
	e := echo.New()
	hello.Init(e)

	e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))


}

func saveUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
