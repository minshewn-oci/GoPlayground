package adapters

import (
	"GoPlayground/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

type IGinWrapper interface {
	Group(path string, handlers ...gin.HandlerFunc) IGinWrapper
	GET(path string, handlerFunc ...gin.HandlerFunc)
}

type GinWrapper struct {
	engine      *gin.Engine
	routerGroup interfaces.RouterGroup
	wrapper     IGinWrapper
}

func NewGinWrapperWithEngine(engine *gin.Engine) *GinWrapper {
	wrapper := new(GinWrapper)
	wrapper.engine = engine
	return wrapper
}

func NewGinWrapperWithRouterGroup(routerGroup interfaces.RouterGroup) *GinWrapper {
	wrapper := new(GinWrapper)
	wrapper.routerGroup = routerGroup
	return wrapper
}

func (w *GinWrapper) Group(path string, handlers ...gin.HandlerFunc) IGinWrapper {
	if w.engine != nil {
		group := w.engine.Group(path, handlers...)
		wrapper := NewGinWrapperWithRouterGroup(group)
		return wrapper
	} else {
		group := w.routerGroup.Group(path, handlers...)
		wrapper := NewGinWrapperWithRouterGroup(group)
		return wrapper
	}
}

func (w *GinWrapper) GET(path string, handlers ...gin.HandlerFunc) {
	if w.engine != nil {
		w.engine.GET(path, handlers...)
	} else {
		w.routerGroup.GET(path, handlers...)
	}
}
