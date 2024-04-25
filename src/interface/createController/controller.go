package createController

import (
	"go-phone-book/src/application/createService"
)

type (
	Controller struct {
		service createService.Service
	}
)

func NewController(service createService.Service) *Controller {
	return &Controller{
		service: service,
	}
}
