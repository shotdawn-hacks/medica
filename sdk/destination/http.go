package destination

import (
	"medview/sdk/client"
)

type Config struct {
	ID      string `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"hostname" bson:"hostname"`
	Port    string `json:"port" bson:"port"`
}

type Destination struct {
	Config *Config
	Client client.Client
	Up     bool
}

func newDestination(cfg *Config) *Destination {
	d := Destination{
		Config: cfg,
	}

	return &d
}

func (r *Destination) Start() {
	err := r.Client.Start()
	if err != nil {
		panic(err)
	}
}

func (r *Destination) HealthCheck() {

}
