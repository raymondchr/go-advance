package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// M is map
type M map[string]interface{}

// ActionHome uses http package to handle the page
var ActionHome = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is ActionHome"))
})

// ActionAbout uses echo WrapHandler to handle the page
var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("this is action about"))
		},
	),
)

// ActionIndex is the basic handler
var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/action", ActionAbout)
	r.Static("/static", "assets")

	r.Start(":9000")
}
