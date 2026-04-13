package hertz

import (
	"context"
	"goravel/db"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Start() {
	h := server.Default(server.WithHostPorts(":9002"))
	h.GET("/", startHandler)
	h.Spin()
}

func startHandler(c context.Context, ctx *app.RequestContext) {
	db.DelayedQuery(c, "Hertz")

	ctx.JSON(http.StatusOK, utils.H{
		"message": "ok",
	})
}
