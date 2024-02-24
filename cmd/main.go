package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lexyu-golang-project-collection/go-templ-project-structure-template/handler"
)

func main() {
	fmt.Println("Hello There!")

	userHandler := handler.UserHandler{}

	app := echo.New()
	app.Use(withUser)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	app.GET("/user", userHandler.UserHandlerShow)

	app.Logger.Fatal(app.Start(":3001"))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// c.Set("user", "tester@gmail.com")
		ctx := context.WithValue(c.Request().Context(), "user", "tester@gmail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
