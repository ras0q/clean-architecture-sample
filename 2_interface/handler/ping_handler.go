package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingHandler interface {
	Ping(c echo.Context) error
}

type pingHandler struct{}

func NewPingHandler() PingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
