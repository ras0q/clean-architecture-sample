package repository

import (
	domain "github.com/Ras96/clean-architecture-sample/0_domain"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/Ras96/clean-architecture-sample/2_interface/repository/model"
	"github.com/google/uuid"
)

type userRepository struct {
	database.SQLHandler
}

func NewUserRepository(sql database.SQLHandler) repository.UserRepository {
	return &userRepository{SQLHandler: sql}
}

// userRepository<struct>がrepository.UserRepository<interface>を満たすようにメソッドを定義する
func (ur *userRepository) FindAll() ([]*domain.User, error) {
	users := make([]*domain.User, 0)
	if err := ur.Find(&users).Error(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	user := domain.NewUser(id, "", "")
	if err := ur.First(&user).Error(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Register(user *repository.RegisteredUser) error {
	newUser := model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	if err := ur.Create(&newUser).Error(); err != nil {
		return err
	}

	return nil
}
