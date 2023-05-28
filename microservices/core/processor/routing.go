package processor

import (
	"medica/microservices/core/api/private"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Core) newAPI() *gin.Engine {
	router := gin.New()

	logger, _ := zap.NewProduction()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	//
	// PRIVATE
	//

	router.POST("/upload", private.Upload)

	router.GET(HTTPHealth, private.Health)
	router.GET("/ping", private.Ping)

	return router
}
