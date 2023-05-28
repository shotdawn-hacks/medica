package processor

import (
	"medica/microservices/core/api/private"
	"medica/sdk/destination"
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

	router.POST(HTTPUpload, r.SetDestiantion(destination.DestinationAnalyzer), private.Upload)

	router.POST(HTTPRegister, SetCore(r), private.Register)
	router.GET(HTTPHealth, private.Health)
	router.GET("/ping", private.Ping)

	return router
}
