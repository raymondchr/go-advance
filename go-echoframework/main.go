package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// M is map
type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}
