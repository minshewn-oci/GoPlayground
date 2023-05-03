package router

import (
	"GoPlayground/internal/factory"
	"GoPlayground/pkg/adapters"
	"github.com/gin-gonic/gin"
)

type Registrar struct {
	HandlerFactory factory.IHandlerFactory
	RouterGroup    adapters.IGinWrapper
}

func (registrar *Registrar) RegisterGroup() {
	group := registrar.RouterGroup.Group("v1")
	registrar.registerEndpoints(group)
}

func (registrar *Registrar) registerEndpoints(group adapters.IGinWrapper) {
	// swagger:route GET /api/v1/ping monitor ping
	//
	// Returns a pong response if the app is running.
	//
	//   Produces:
	//	 - application/json
	//
	//   Schemes: http, https
	//
	//   Responses:
	//     200: pingResponse

	var handler gin.HandlerFunc
	handler = func(c *gin.Context) {
		registrar.HandlerFactory.Get("ping").Process(c)
	}
	group.GET("ping", handler)
}
