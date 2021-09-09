package handler

import (
	"net/http"

	usecase "github.com/Ras96/clean-architecture-sample/1_usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	uc usecase.UserUsecase //TODO: userServiceの方がいいかも
}

//TODO: struct返すのかここ
func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	return c.NoContent(http.StatusOK) //TODO
}

func (h *UserHandler) GetByID(c echo.Context) error {
	return c.NoContent(http.StatusOK) //TODO
}

func (h *UserHandler) Register(c echo.Context) error {
	return c.NoContent(http.StatusOK) //TODO
}
