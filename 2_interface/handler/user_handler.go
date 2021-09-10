package handler

import (
	"errors"
	"net/http"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/1_usecase"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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

type UserHandler struct {
	uc usecase.UserService
}

//TODO: interface返すためにはinfrastructure層のserviceをinterfaceで定義しておく必要がありそう
func NewUserHandler(uc usecase.UserService) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetAll(c echo.Context) error {
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

func (h *UserHandler) GetByID(c echo.Context) error {
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

func (h *UserHandler) Register(c echo.Context) error {
	req := RegisterReq{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.uc.Register(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
