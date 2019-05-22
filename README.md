# Fasthttp Router Middleware

[![CircleCI](https://circleci.com/gh/roger-russel/fasthttp-router-middleware.svg?style=shield)](https://circleci.com/gh/roger-russel/fasthttp-router-middleware) [![codecov](https://codecov.io/gh/roger-russel/fasthttp-router-middleware/branch/master/graph/badge.svg)](https://codecov.io/gh/roger-russel/fasthttp-router-middleware) [![Software License](https://img.shields.io/badge/license-Apache-brightgreen.svg?style=flat-square)](LICENSE.md)

## Usage

There is two usages of this Ondemand and a created group.

### Ondemand

Ondemand you just put yours middlewares per route like the example bellow.
There is a working example of this [here](./doc/examples/ondemand) with its tests.

```go

import "github.com/roger-russel/fasthttp-router-middleware/pkg/middleware"

func exampleAuthFunc(ctx *fasthttp.RequestCtx) bool { ... }
func exampleRuleFunc(ctx *fasthttp.RequestCtx) bool { ... }
func exampleRequestHandler(ctx *fasthttp.RequestCtx) { ... }

func main() {

  ...

	router := router.New()
	router.GET("/", exampleRequestHandler)
	router.GET("/protected", middleware.Apply([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc}, exampleRequestHandler))

  ...

}

```

### Middleware Groups

Middleware grous allow you to create a group of meiddlewares and reuse it into many routes that you like.
There is a working example of this [here](./doc/examples/group) with its tests.

```go

import "github.com/roger-russel/fasthttp-router-middleware/pkg/middleware"

func exampleAuthFunc(ctx *fasthttp.RequestCtx) bool { ... }
func exampleRuleFunc(ctx *fasthttp.RequestCtx) bool { ... }
func exampleRequestHandler(ctx *fasthttp.RequestCtx) { ... }

func main() {

  ...

  midGroupAuth = middleware.New([]middleware.Middleware{exampleAuthFunc, exampleRuleFunc})

	router := router.New()
	router.GET("/", exampleRequestHandler)
  router.GET("/protected", midGroupAuth(exampleRequestHandler))

  ...

}

```

## Contribute Guide

Please take a look at [Contribute Guide](./doc/contributing.md).

## Thanks

* @Hanjm, I learn a lot on [his middleware project](https://github.com/hanjm/fasthttpmiddleware) for fasthttp please take a look there too.
