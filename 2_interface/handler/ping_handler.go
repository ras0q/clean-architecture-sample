//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

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
