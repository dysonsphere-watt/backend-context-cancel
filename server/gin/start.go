package gin

import (
	"goravel/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/", startHandler)
	r.Run(":9001")
}

func startHandler(c *gin.Context) {
	db.DelayedQuery(c.Request.Context(), "Gin")

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
