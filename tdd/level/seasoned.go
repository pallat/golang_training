package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello, wrapper("hello"))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type myContext struct {
	echo.Context
	name string
}

func wrapper(name string) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&myContext{c, name})
		}
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! my name is "+c.(*myContext).name)
}
