package processor

import (
	"fmt"
	"medica/sdk/destination"
	"sync"

	"github.com/gin-gonic/gin"
)

type Core struct {
	ID           string                     `json:"_id" bson:"_id"`
	Address      string                     `json:"address" bson:"address"`
	Port         string                     `json:"port" bson:"port"`
	Destinations []*destination.Destination `json:"destinations" bson:"destinations"`
	mu           sync.RWMutex
}

type Config struct {
	Plants *destination.Config
}

func (r *Core) AppendDestiantion(cfg *destination.Config) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.Destinations = append(r.Destinations, destination.NewDestination(cfg))
}

func (r *Core) SetDestiantion(name string) gin.HandlerFunc {
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

func NewDefaultCore(cfgs ...destination.Config) *Core {
	c := &Core{
		Address: "",
		Port:    "9010",
	}

	for _, cfg := range cfgs {
		c.Destinations = append(c.Destinations, destination.NewDestination(&cfg))
	}

	return c
}

func (r *Core) init() {
	for _, dest := range r.Destinations {
		if err := dest.Start(); err != nil {
			panic(err)
		}
	}
}

func (r *Core) Start() {
	r.init()

	api := r.newAPI()

	api.Run(":" + r.Port)
}
