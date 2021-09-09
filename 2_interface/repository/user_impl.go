package repository

import (
	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	sql database.SQLHandler
}

func NewUserRepository(sql database.SQLHandler) repository.UserRepository {
	return &UserRepository{sql}
}

// UserRepository(構造体)がrepository.UserRepository(インターフェース)を満たすためにメソッドを定義する
func (ur *UserRepository) FindAll() ([]*model.User, error) {
	// DBから取得するロジックを書く
	return nil, nil //TODO
}

func (ur *UserRepository) FindByID(id uuid.UUID) (*model.User, error) {
	return nil, nil //TODO
}

func (ur *UserRepository) Create(*model.User) error {
	return nil //TODO
}
