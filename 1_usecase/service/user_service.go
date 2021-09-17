//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package service

import (
	"fmt"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/gofrs/uuid"
)

type UserService interface {
	GetAll() ([]*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	Register(user *repository.RegisteredUser) error
}

type userService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

// userService(構造体)がUserService(インターフェース)を満たすためにメソッドを定義する
func (s *userService) GetAll() ([]*model.User, error) {
	users, err := s.user.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (s *userService) GetByID(id uuid.UUID) (*model.User, error) {
	user, err := s.user.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (s *userService) Register(user *repository.RegisteredUser) error {
	if err := s.user.Register(user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
