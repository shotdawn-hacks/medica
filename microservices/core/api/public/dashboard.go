package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"dasboard":"cool"}`)
}
