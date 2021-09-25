package registry

import (
	controller "github.com/conzmr/academy-go-q32021/interface/controllers"
)

type registry struct {
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewCoachController()
}
