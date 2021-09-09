package usecase

import (
	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/gofrs/uuid"
)

//TODO: 別のファイルに置きたい
type UserUsecase interface {
	GetAll() ([]*model.User, error) //TODO:名前を考える
	GetByID(id uuid.UUID) (*model.User, error)
	Register(user *model.User) error //TODO:model.Userとは別の構造体を用意するべき
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

// userUsecase(構造体)がUserUsecase(インターフェース)を満たすためにメソッドを定義する
func (uc *userUsecase) GetAll() ([]*model.User, error) {
	return nil, nil //TODO
}

func (uc *userUsecase) GetByID(id uuid.UUID) (*model.User, error) {
	return nil, nil //TODO
}

func (uc *userUsecase) Register(user *model.User) error {
	return nil //TODO
}
