//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package repository

import (
	"github.com/google/uuid"
	domain "github.com/ras0q/clean-architecture-sample/0_domain"
)

type RegisteredUser struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	Register(user *RegisteredUser) error
}
