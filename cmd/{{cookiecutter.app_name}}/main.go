package main

import (
	"{{cookiecutter.app_name}}/pkg/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Routes
	e.GET("/", view.Index)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
