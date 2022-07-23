package controllers

import (
	"net/http"

	"github.com/api-server-proxy/internal/router"
)

type storageController struct{}

func NewStorageController() *storageController {
	return &storageController{}
}

func (s *storageController) RegisterEndpoints(router router.Router) {
	group := router.Group("storage")

	group.Post("/key/{key}", s.create)

}

func (s *storageController) create(http.ResponseWriter, *http.Request) {

}
