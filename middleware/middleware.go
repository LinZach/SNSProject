package middleware

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type midleware_t func(ctx *fasthttp.RequestCtx) error

var MiddlewareList = []midleware_t{
	CorsMiddleware,
	CheckTokenMiddleware,
}

func User(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		for _, r := range MiddlewareList {
			if err := r(ctx); err != nil {
				fmt.Print(err)
				return
			}
		}
		next(ctx)
	})
}
