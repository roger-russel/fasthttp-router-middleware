package middleware

import "github.com/valyala/fasthttp"

// New Middleware group.
func New(middlewares []Middleware) func(fasthttp.RequestHandler) fasthttp.RequestHandler {

	return func(currentRoute fasthttp.RequestHandler) fasthttp.RequestHandler {

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

}
