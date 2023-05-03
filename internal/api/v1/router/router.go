package router

import (
	"GoPlayground/internal/factory"
	"GoPlayground/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

type Registrar struct {
	HandlerFactory factory.IHandlerFactory
}

func (registrar *Registrar) RegisterGroup(group interfaces.RouterGroup) interfaces.RouterGroup {
	group = group.Group("v1")
	registrar.registerEndpoints(group)
	return group
}

func (registrar *Registrar) registerEndpoints(group interfaces.RouterGroup) {
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
