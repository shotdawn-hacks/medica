package private

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
