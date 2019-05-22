package main

import (
	"bytes"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/roger-russel/fasthttp-router-middleware/pkg/middleware"
	"github.com/valyala/fasthttp"
)

func exampleAuthFunc(ctx *fasthttp.RequestCtx) bool {

	if bytes.HasPrefix(ctx.Path(), []byte("/unauthorized")) {
		ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
		return false
	}

	return true
}

func exampleRuleFunc(ctx *fasthttp.RequestCtx) bool {

	if bytes.HasPrefix(ctx.Path(), []byte("/forbidden")) {
		ctx.Response.SetStatusCode(fasthttp.StatusForbidden)
		return false
	}

	return true
}

func exampleRequestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}

func main() {

	midAccessGroup := middleware.New([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc})

	router := router.New()
	router.GET("/", exampleRequestHandler)
	router.GET("/unauthorized", midAccessGroup(exampleRequestHandler))
	router.GET("/forbidden", midAccessGroup(exampleRequestHandler))
	router.GET("/authorized", midAccessGroup(exampleRequestHandler))

	if err := fasthttp.ListenAndServe(":8100", router.Handler); err != nil {
		fmt.Println("Error in ListenAndServe: ", err)
	}

}
