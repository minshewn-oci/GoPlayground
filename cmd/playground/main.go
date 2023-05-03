// Package classification Playground API.
//
// The purpose of this API is to provide an example of how a Golang API could be constructed.
//
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"GoPlayground/internal/api/router"
	"GoPlayground/internal/factory"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteRegistrar[T any] interface {
	RegisterGroup(router T) T
}

func main() {
	r := gin.Default()
	var handlerFactory factory.IHandlerFactory
	handlerFactory = new(factory.HandlerFactory)
	apiRegistrar := router.Registrar{HandlerFactory: handlerFactory}
	apiRegistrar.RegisterGroup(r)
	r.Static("/files/openapi-spec/", "api/openapi-spec")
	r.Static("/docs", "docs")
	if err := r.Run(":3000"); err != nil {
		fmt.Println(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
