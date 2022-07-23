package controllers

import "github.com/api-server-proxy/internal/router"

type controllerManager struct {
	router router.Router
}

func (c *controllerManager) RegisterEndpoints() {

	_ = c.router.Group("v1")
}
