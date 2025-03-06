package main

import (
	"fmt"

	"{{cookiecutter.app_name}}/pkg/view"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	config.AddDriver(json.Driver)
	env := config.GetEnv("{{cookiecutter.app_name|upper}}_ENV", "dev")
	err := config.LoadFiles("config/" + env + ".json")
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		panic(err)
	}
	err = config.LoadFiles("config/secrets.json")
	if err != nil {
		fmt.Printf("error loading secrets: %v\n", err)
		fmt.Println("if the task is running in ECS, the secrets file is not needed")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/marco"
		},
	}))
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Routes
	e.GET("/", view.Index)
	e.GET("/marco", view.Marco)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
