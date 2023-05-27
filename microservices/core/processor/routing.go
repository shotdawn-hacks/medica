package processor

import (
	"medica/microservices/core/api/private"

	"github.com/gin-gonic/gin"
)

func (r *Core) newAPI() *gin.Engine {
	router := gin.New()
	// publicRouter := router.Group("/api/v1")

	//
	// PRIVATE
	//
	router.POST("/upload", private.Upload)

	router.GET(HTTPHealth, private.Health)
	router.GET("/ping", private.Ping)

	return router
}
