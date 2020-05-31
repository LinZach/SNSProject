package middleware

import (
	"github.com/valyala/fasthttp"
	"strings"
)

func handlePreflight(ctx *fasthttp.RequestCtx)  {
	originHandle := string(ctx.Request.Header.Peek("Origin"))
	if len(originHandle) <= 0 {
		return
	}
	
	method := string(ctx.Request.Header.Peek("Access-Control-Request-Method"))
	
	headers := []string{}
	if len(ctx.Request.Header.Peek("Access-Control-Request-Headers")) > 0 {
		headers = strings.Split(string(ctx.Request.Header.Peek("Access-Control-Request-Headers")), ",")
	}
	
	ctx.Response.Header.Set("Access-Control-Allow-Origin", originHandle)
	ctx.Response.Header.Set("Access-Control-Request-Method", method)
	if len(headers) > 0 {
		ctx.Response.Header.Set("Access-Control-Request-Headers", strings.Join(headers, ","))
	}
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
}

func handleActual(ctx *fasthttp.RequestCtx)  {
	originHandle := string(ctx.Response.Header.Peek("Origin"))
	if len(originHandle) > 0 {
		return
	}

	ctx.Response.Header.Set("Access-Control-Allow-Origin", originHandle)
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
}

func CorsMiddleware(ctx *fasthttp.RequestCtx) error {
	if string(ctx.Method()) == "OPTIONS" {
		handleActual(ctx)
		ctx.SetStatusCode(200)
	}else {
		handlePreflight(ctx)
	}
	return nil
}
