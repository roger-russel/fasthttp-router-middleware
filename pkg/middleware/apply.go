package middleware

import "github.com/valyala/fasthttp"

// Middleware must return true for continue or false to stop on it.
type Middleware func(ctx *fasthttp.RequestCtx) bool

// Apply Middleware List into route.
func Apply(middlewares []Middleware, currentRoute fasthttp.RequestHandler) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {

		var applied bool

		for _, m := range middlewares {

			applied = m(ctx)

			if !applied {
				return
			}

		}

		currentRoute(ctx)

	}

}
