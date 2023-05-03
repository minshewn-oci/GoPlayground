package factory

import (
	"GoPlayground/internal/api/v1/handlers/ping"
	"fmt"
)

type IHandlerFactory interface {
	Get(path string) ping.IHandler
}

type HandlerFactory struct{}

func (f *HandlerFactory) Get(path string) ping.IHandler {
	switch path {
	case "ping":
		return new(ping.PingHandler)
	default:
		panic(fmt.Sprintf("Unable to find handler with path %s", path))
	}
}
