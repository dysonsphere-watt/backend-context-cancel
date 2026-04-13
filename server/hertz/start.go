package hertz

import (
	"context"
	"fmt"
	"goravel/app/facades"
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
	facades.Orm().WithContext(c).Query().Exec("WAITFOR DELAY '0:00:10'")
	fmt.Print("\n[Hertz] Wait DB call done\n\n")

	ctx.JSON(http.StatusOK, utils.H{
		"message": "ok",
	})
}
