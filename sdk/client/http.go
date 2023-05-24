package client

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPClient struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Hostname string `json:"hostname" bson:"hostname"`
	Port     int    `json:"port" bson:"port"`
}

func (r *HTTPClient) send(ctx *gin.Context) {
	c := &http.Client{}

	req := ctx.Copy().Request

	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}

	ctx.Status(resp.StatusCode)
}
