package client

import "github.com/gin-gonic/gin"

type Client interface {
	Start() error
	send(ctx *gin.Context)
	health() uint16
}
