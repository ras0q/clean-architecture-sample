package repository

import (
	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/gofrs/uuid"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	Create(user *model.User) error
}
