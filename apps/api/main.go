package main

import (
	"api/libs/api/greeting"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"greeting": greeting.Greeting("Hello")})
	})

	e.Logger.Fatal(e.Start(":3333"))
}
