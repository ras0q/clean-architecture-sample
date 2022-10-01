//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package handler

import (
	"errors"
	"net/http"

	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/Ras96/clean-architecture-sample/1_usecase/service"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetAll(c Context) error
	GetByID(c Context) error
	Register(c Context) error
}

type userHandler struct {
	srv service.UserService
}

func NewUserHandler(srv service.UserService) UserHandler {
	return &userHandler{srv}
}

type UserRes struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserDetailRes struct {
	UserRes
	Email string `json:"email"`
}

type RegisterReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// userHandler<struct>がUserHandler<interface>を満たすようにメソッドを定義する

// GetAll GET /users
func (h *userHandler) GetAll(c Context) error {
	users, err := h.srv.GetAll()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := make([]*UserRes, 0, len(users))
	for _, v := range users {
		res = append(res, &UserRes{
			ID:   v.ID(),
			Name: v.Name(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// GetByID GET /users/:id
func (h *userHandler) GetByID(c Context) error {
	idstr := c.Param("id")
	id, err := uuid.Parse(idstr)
	if err != nil || id == uuid.Nil {
		return c.JSON(http.StatusBadRequest, err.Error()) // invalid uuid
	}

	user, err := h.srv.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := &UserDetailRes{
		UserRes: UserRes{
			ID:   user.ID(),
			Name: user.Name(),
		},
		Email: user.Email(),
	}

	return c.JSON(http.StatusOK, res)
}

// Pegister POST /users
func (h *userHandler) Register(c Context) error {
	req := RegisterReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := repository.RegisteredUser{
		ID:    random.UUID(),
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.srv.Register(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
