package processor

import (
	"medica/sdk/destination"
)

type Core struct {
	ID           string                     `json:"_id" bson:"_id"`
	Address      string                     `json:"address" bson:"address"`
	Port         string                     `json:"port" bson:"port"`
	Destinations []*destination.Destination `json:"destinations" bson:"destinations"`
}

type Config struct {
	Plants *destination.Config
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
