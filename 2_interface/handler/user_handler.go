package handler

import (
	"errors"
	"net/http"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/1_usecase/service"
	"github.com/gofrs/uuid"
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
	uc service.UserService
}

//TODO: interface返すためにはinfrastructure層のserviceをinterfaceで定義しておく必要がありそう
func NewUserHandler(uc service.UserService) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetAll(c Context) error {
	users, err := h.uc.GetAll()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

func (h *UserHandler) GetByID(c Context) error {
	idstr := c.Param("id")
	id, err := uuid.FromString(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()) // invalid uuid
	}

	user, err := h.uc.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

func (h *UserHandler) Register(c Context) error {
	req := RegisterReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.uc.Register(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
