package router

import (
	"GoPlayground/internal/api/v1/router"
	"GoPlayground/internal/factory"
	"GoPlayground/pkg/adapters"
)

type Registrar struct {
	HandlerFactory factory.IHandlerFactory
	RouterGroup    adapters.IGinWrapper
}

func (registrar *Registrar) RegisterGroup() {
	api := registrar.RouterGroup.Group("api")
	wrapper := adapters.NewGinWrapperWithRouterGroup(api)
	v1Registrar := router.Registrar{HandlerFactory: registrar.HandlerFactory, RouterGroup: wrapper}
	v1Registrar.RegisterGroup()
}
