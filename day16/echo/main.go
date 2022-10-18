package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID string `param:"id"`
}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) foo(e *echo.Echo) {
	fmt.Println("Bar")
	e.Logger.Info("Info message")
}

func main() {
	e := echo.New()


	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/:id", func(c echo.Context) error {
		cc := c.(*CustomContext)
		var user User

		err := cc.Bind(&user); if err != nil {
			return cc.String(http.StatusBadRequest, "bad request")
		}
		cc.foo(e)
		return cc.String(http.StatusOK, user.ID)
	})

	e.Logger.Fatal(e.Start(":8081"))
}