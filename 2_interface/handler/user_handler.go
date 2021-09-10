package handler

import (
	"errors"
	"net/http"

	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	usecase "github.com/Ras96/clean-architecture-sample/1_usecase"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Register(c echo.Context) error
}

type userHandler struct {
	uc usecase.UserService
}

func NewUserHandler(uc usecase.UserService) UserHandler {
	return &userHandler{uc}
}

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserDetail struct {
	User
	Email string `json:"email"`
}

type RegisterReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *userHandler) GetAll(c echo.Context) error {
	users, err := h.uc.GetAll()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := make([]*User, 0, len(users))
	for _, v := range users {
		res = append(res, &User{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandler) GetByID(c echo.Context) error {
	idstr := c.Param("id")
	id, err := uuid.FromString(idstr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()) // invalid uuid
	}

	user, err := h.uc.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := &UserDetail{
		User: User{
			ID:   user.ID,
			Name: user.Name,
		},
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *userHandler) Register(c echo.Context) error {
	req := RegisterReq{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := repository.RegisteredUser{
		ID:    uuid.Must(uuid.NewV4()),
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.uc.Register(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
