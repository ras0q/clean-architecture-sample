package handler

import (
	"net/http"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c Context) error {
	return c.String(http.StatusOK, "pong!")
}
