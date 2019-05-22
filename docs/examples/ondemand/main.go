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

	router := router.New()
	router.GET("/", exampleRequestHandler)
	router.GET("/unauthorized", middleware.Apply([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc}, exampleRequestHandler))
	router.GET("/forbidden", middleware.Apply([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc}, exampleRequestHandler))
	router.GET("/authorized", middleware.Apply([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc}, exampleRequestHandler))

	if err := fasthttp.ListenAndServe(":8101", router.Handler); err != nil {
		fmt.Println("Error in ListenAndServe: ", err)
	}

}
