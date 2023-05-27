package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
