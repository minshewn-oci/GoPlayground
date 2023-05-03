package ping

import (
	"net/http"
)

type Context interface {
	JSON(code int, object any)
}

// A PingResponse is a response to a ping.
// swagger:response pingResponse
type PingResponse struct {
	// The message returned
	Message string
}

type IHandler interface {
	Process(c Context)
}

type PingHandler struct{}

func (h *PingHandler) Process(c Context) {
	response := PingResponse{Message: "pong"}
	c.JSON(http.StatusOK, response)
}
