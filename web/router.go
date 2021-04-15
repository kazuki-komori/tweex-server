package web

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")

	v1.GET("/health", health)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func health(c echo.Context) error {
	type health struct {
		Status string `json:"status"`
	}
	res := new(health)
	res.Status = "OK"
	return c.JSON(http.StatusOK, res)
}
