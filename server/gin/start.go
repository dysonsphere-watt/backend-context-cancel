package gin

import (
	"fmt"
	"goravel/app/facades"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/", startHandler)
	r.Run(":9001")
}

func startHandler(c *gin.Context) {
	facades.Orm().WithContext(c.Request.Context()).Query().Exec("WAITFOR DELAY '0:00:10'")
	fmt.Print("\n[Gin] Wait DB call done\n\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
