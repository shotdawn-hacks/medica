package client

import (
	"fmt"
	"net/http"
)

type Base struct {
	Client *http.Client
}

func (r *Base) Start() bool {
	r.init()

	return true
}

func (r *Base) init() {
	r.Client = new(http.Client)
}

func (r *Base) Get(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodGet {
		return nil, fmt.Errorf("expected GET request, got %s", req.Method)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Base) Post(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodPost {
		return nil, fmt.Errorf("expected POST request, got %s", req.Method)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
