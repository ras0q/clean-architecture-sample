package usecase

import (
	"fmt"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/gofrs/uuid"
)

//TODO: 別のファイルに置きたい
type UserService interface {
	GetAll() ([]*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	Register(user *model.User) error //TODO:model.Userとは別の構造体を用意するべき
}

type userSerUserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userSerUserService{repo}
}

// userSerUserService(構造体)がUserService(インターフェース)を満たすためにメソッドを定義する
// GET /users
func (uc *userSerUserService) GetAll() ([]*model.User, error) {
	users, err := uc.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

// GET /users/:id
func (uc *userSerUserService) GetByID(id uuid.UUID) (*model.User, error) {
	user, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

// POST /users/
func (uc *userSerUserService) Register(user *model.User) error {
	if err := uc.repo.Register(user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
