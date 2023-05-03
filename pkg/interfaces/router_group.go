package interfaces

import "github.com/gin-gonic/gin"

type RouterGroup interface {
	Group(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup
	GET(path string, handlerFunc ...gin.HandlerFunc) gin.IRoutes
}
