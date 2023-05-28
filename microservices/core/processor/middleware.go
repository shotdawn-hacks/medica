package processor

import (
	"medica/sdk/destination"

	"github.com/gin-gonic/gin"
)

func SetCore(core destination.Setter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("core", core)
		ctx.Next()
	}
}
