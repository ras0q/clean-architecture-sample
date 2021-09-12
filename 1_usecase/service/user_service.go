package service

import (
	"fmt"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/gofrs/uuid"
)

type UserService interface {
	GetAll() ([]*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	Register(user *repository.RegisteredUser) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

// userService(構造体)がUserService(インターフェース)を満たすためにメソッドを定義する
func (uc *userService) GetAll() ([]*model.User, error) {
	users, err := uc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (uc *userService) GetByID(id uuid.UUID) (*model.User, error) {
	user, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (uc *userService) Register(user *repository.RegisteredUser) error {
	if err := uc.repo.Register(user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
