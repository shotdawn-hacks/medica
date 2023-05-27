package destination

import (
	"fmt"
	"medica/sdk/client"
)

type Destination struct {
	Config *Config
	Base   client.Base
	Up     bool
}

func NewDestination(cfg *Config) *Destination {
	d := Destination{
		Config: cfg,
	}

	return &d
}

func (r *Destination) Start() error {
	ok := r.Base.Start()
	if !ok {
		return fmt.Errorf("Base failed to start")
	}

	return nil
}
