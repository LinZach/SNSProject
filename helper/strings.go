package helper

import (
	"github.com/valyala/fasthttp"
)
import "github.com/wI2L/jettison"

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}


func Print(ctx *fasthttp.RequestCtx, status string, data interface{}) {
	ctx.SetStatusCode(200)
	res := Response{status, data}

	bytes, err := jettison.Marshal(res)

	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}

	ctx.SetBody(bytes)

}