package private

import (
	"encoding/json"
	"fmt"
	"medica/sdk/destination"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func Register(ctx *gin.Context) {
	var destCfg destination.Config

	json.NewDecoder(ctx.Request.Body).Decode(&destCfg)

	core, ok := ctx.Get("core")
	if !ok {
		ctx.AbortWithError(http.StatusOK, fmt.Errorf("core is not exists"))
	}

	coreSetter, ok := core.(destination.Setter)
	if !ok {
		ctx.AbortWithError(http.StatusOK, fmt.Errorf("core is wrong interface"))
	}

	coreSetter.AppendDestiantion(&destCfg)
}
