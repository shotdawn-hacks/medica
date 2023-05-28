package processor

import (
	"fmt"
	"medica/microservices/api-gateway/api/private"
	"medica/microservices/api-gateway/api/public"
	"medica/sdk/destination"
	"medica/sdk/shared"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Gateway struct {
	ID           string                     `json:"_id" bson:"_id"`
	Address      string                     `json:"address" bson:"address"`
	Port         string                     `json:"port" bson:"port"`
	Destinations []*destination.Destination `json:"destinations" bson:"destinations"`
}

func NewDefaultGateway(cfgs ...destination.Config) *Gateway {
	c := &Gateway{
		Address: "",
		Port:    "9000",
	}
	for _, cfg := range cfgs {
		c.Destinations = append(c.Destinations, destination.NewDestination(&cfg))
	}
	return c
}

func (r *Gateway) init() {
	for _, dest := range r.Destinations {
		if err := dest.Start(); err != nil {
			panic(err)
		}
	}
}

func (r *Gateway) Start() {
	r.init()

	api := r.newAPI()

	api.Run(":" + r.Port)
}

func (r *Gateway) SetDestiantion(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, dst := range r.Destinations {
			if dst.Config.Name == name {
				ctx.Set(fmt.Sprintf("dst-%s", name), dst)
				ctx.Next()
				break
			}
		}

	}
}

func (r *Gateway) newAPI() *gin.Engine {
	router := gin.New()

	logger, _ := zap.NewProduction()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	publicRouter := router.Group("/api/v1")

	//
	// CORS
	//

	router.Use(shared.SetDefaultCors())

	//
	// PUBLIC
	//

	publicRouter.Use(r.SetDestiantion("core"))

	publicRouter.GET(HTTPDashboard, public.Dashboard)

	publicRouter.POST("/upload", public.Upload)

	//
	// PRIVATE
	//

	router.GET(HTTPHealth, private.Health)

	return router
}
