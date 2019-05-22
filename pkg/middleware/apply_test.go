package middleware

import (
	"bytes"
	"testing"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func TestApplySingle(t *testing.T) {

	port := getAvaliablePort()

	var url = "http://127.0.0.1:" + port

	exampleAuthFunc := func(ctx *fasthttp.RequestCtx) bool {

		if bytes.HasPrefix(ctx.Path(), []byte("/unauthorized")) {
			ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
			return false
		}
		return true
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("hello")
	}

	router := router.New()
	router.GET("/", requestHandler)
	router.GET("/authorized", Apply([]Middleware{exampleAuthFunc}, requestHandler))
	router.GET("/unauthorized", Apply([]Middleware{exampleAuthFunc}, requestHandler))

	doneChan := make(chan struct{})
	go func() {
		fasthttp.ListenAndServe(":"+port, router.Handler)
	}()

	delay()

	go func() {

		var resp []byte

		code, _, _ := fasthttp.Get(resp, url)
		if code != 200 {
			t.Errorf(errorCode(200, code))
		}

		code, _, _ = fasthttp.Get(resp, url+"/unauthorized")
		if code != 401 {
			t.Errorf(errorCode(401, code))
		}

		code, _, _ = fasthttp.Get(resp, url+"/authorized")
		if code != 200 {
			t.Errorf(errorCode(401, code))
		}

		doneChan <- struct{}{}
	}()

	<-doneChan
}

func TestApplyMultiples(t *testing.T) {

	port := getAvaliablePort()

	var url = "http://127.0.0.1:" + port

	exampleAuthFunc := func(ctx *fasthttp.RequestCtx) bool {

		if bytes.HasPrefix(ctx.Path(), []byte("/unauthorized")) {
			ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
			return false
		}
		return true
	}

	exampleRoleFunc := func(ctx *fasthttp.RequestCtx) bool {

		if bytes.HasPrefix(ctx.Path(), []byte("/forbidden")) {
			ctx.Response.SetStatusCode(fasthttp.StatusForbidden)
			return false
		}
		return true
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("hello")
	}

	router := router.New()
	router.GET("/", requestHandler)
	router.GET("/authorized", Apply([]Middleware{exampleAuthFunc, exampleRoleFunc}, requestHandler))
	router.GET("/unauthorized", Apply([]Middleware{exampleAuthFunc, exampleRoleFunc}, requestHandler))
	router.GET("/forbidden", Apply([]Middleware{exampleAuthFunc, exampleRoleFunc}, requestHandler))

	doneChan := make(chan struct{})
	go func() {
		fasthttp.ListenAndServe(":"+port, router.Handler)
	}()

	delay()

	go func() {

		var resp []byte

		code, _, _ := fasthttp.Get(resp, url)
		if code != 200 {
			t.Errorf(errorCode(200, code))
		}

		code, _, _ = fasthttp.Get(resp, url+"/unauthorized")
		if code != 401 {
			t.Errorf(errorCode(401, code))
		}

		code, _, _ = fasthttp.Get(resp, url+"/authorized")
		if code != 200 {
			t.Errorf(errorCode(200, code))
		}

		code, _, _ = fasthttp.Get(resp, url+"/forbidden")
		if code != 403 {
			t.Errorf(errorCode(403, code))
		}

		doneChan <- struct{}{}
	}()

	<-doneChan
}
