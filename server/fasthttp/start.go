package fasthttp

import (
	"encoding/json"
	"goravel/db"
	"net/http"

	"github.com/valyala/fasthttp"
)

func Start() {
	fasthttp.ListenAndServe(":9003", startHandler)
}

func startHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != "GET" || string(ctx.Path()) != "/" {
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	db.DelayedQuery(ctx, "FastHTTP")

	body, _ := json.Marshal(map[string]string{"message": "ok"})
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(body)
}
