package handler

import (
	"errors"
	"net/http"

	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/1_usecase/service"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetAll(c Context) error
	GetByID(c Context) error
	Register(c Context) error
}

type userHandler struct {
	uc service.UserService
}

func NewUserHandler(uc service.UserService) UserHandler {
	return &userHandler{uc}
}

type userRes struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type userDetailRes struct {
	userRes
	Email string `json:"email"`
}

type registerReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GET /users
func (h *userHandler) GetAll(c Context) error {
	users, err := h.uc.GetAll()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := make([]*userRes, 0, len(users))
	for _, v := range users {
		res = append(res, &userRes{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return c.JSON(http.StatusOK, res)
}

// GET /users/:id
func (h *userHandler) GetByID(c Context) error {
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

	res := &userDetailRes{
		userRes: userRes{
			ID:   user.ID,
			Name: user.Name,
		},
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /users
func (h *userHandler) Register(c Context) error {
	req := registerReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := repository.RegisteredUser{
		ID:    uuid.Must(uuid.NewV4()),
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.uc.Register(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
