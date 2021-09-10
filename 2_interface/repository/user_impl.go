package repository

import (
	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	database.SQLHandler
}

func NewUserRepository(sql database.SQLHandler) repository.UserRepository {
	return &UserRepository{SQLHandler: sql}
}

// UserRepository(構造体)がrepository.UserRepository(インターフェース)を満たすためにメソッドを定義する
func (ur *UserRepository) FindAll() ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := ur.Find(&users).Error(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) FindByID(id uuid.UUID) (*model.User, error) {
	user := model.User{ID: id}
	if err := ur.First(&user).Error(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Register(user *model.User) error {
	if err := ur.Create(user).Error(); err != nil {
		return err
	}

	return nil
}
