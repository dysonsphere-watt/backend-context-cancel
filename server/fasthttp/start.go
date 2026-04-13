package fasthttp

import (
	"encoding/json"
	"fmt"
	"goravel/app/facades"
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

	facades.Orm().WithContext(ctx).Query().Exec("WAITFOR DELAY '0:00:10'")
	fmt.Print("\n[FastHTTP] Wait DB call done\n\n")

	body, _ := json.Marshal(map[string]string{"message": "ok"})
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(body)
}
