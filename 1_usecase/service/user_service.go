//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package service

import (
	"fmt"

	domain "github.com/Ras96/clean-architecture-sample/0_domain"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/google/uuid"
)

type UserService interface {
	GetAll() ([]*domain.User, error)
	GetByID(id uuid.UUID) (*domain.User, error)
	Register(user *repository.RegisteredUser) error
}

type userService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}

// userService<struct>がUserService<interface>を満たすようにメソッドを定義する
func (s *userService) GetAll() ([]*domain.User, error) {
	users, err := s.user.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (s *userService) GetByID(id uuid.UUID) (*domain.User, error) {
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
