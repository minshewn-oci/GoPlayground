package router

import (
	"GoPlayground/internal/api/v1/router"
	"GoPlayground/internal/factory"
	"GoPlayground/pkg/interfaces"
)

type Registrar struct {
	HandlerFactory factory.IHandlerFactory
}

func (registrar *Registrar) RegisterGroup(engine interfaces.RouterGroup) interfaces.RouterGroup {
	v1Registrar := router.Registrar{HandlerFactory: registrar.HandlerFactory}
	api := engine.Group("api")
	v1Registrar.RegisterGroup(api)
	return api
}
