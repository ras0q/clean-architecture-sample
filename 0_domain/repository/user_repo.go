//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package repository

import (
	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/gofrs/uuid"
)

type RegisteredUser struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	Register(user *RegisteredUser) error
}
