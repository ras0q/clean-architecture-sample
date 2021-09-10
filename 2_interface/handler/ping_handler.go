package handler

import (
	"net/http"
)

type PingHandler interface {
	Ping(c Context) error
}

type pingHandler struct{}

func NewPingHandler() PingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Ping(c Context) error {
	return c.String(http.StatusOK, "pong!")
}
