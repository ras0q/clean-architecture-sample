package service

import (
	"fmt"
	"testing"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository/mock_repository"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	t.Parallel()
	type args struct {
		repo repository.UserRepository
	}
	tests := []struct {
		name  string
		args  args
		want  UserService
		setup func(args args, want UserService)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Assertion
			assert.Equalf(t, tt.want, NewUserService(tt.args.repo), "NewUserService(%v)", tt.args.repo)
		})
	}
}

func Test_userService_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		want      []*model.User
		setup     func(repo *mock_repository.MockUserRepository, want []*model.User)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.want)
			uc := NewUserService(repo)
			// Assertion
			got, err := uc.GetAll()
			tt.assertion(t, err, fmt.Sprintf("userService.GetAll()"))
			assert.Equalf(t, tt.want, got, "userService.GetAll()")
		})
	}
}

func Test_userService_GetByID(t *testing.T) {
	t.Parallel()
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name      string
		args      args
		want      *model.User
		setup     func(repo *mock_repository.MockUserRepository, args args, want *model.User)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.args, tt.want)
			uc := NewUserService(repo)
			// Assertion
			got, err := uc.GetByID(tt.args.id)
			tt.assertion(t, err, fmt.Sprintf("userService.GetByID(%v)", tt.args.id))
			assert.Equalf(t, tt.want, got, "userService.GetByID(%v)", tt.args.id)
		})
	}
}

func Test_userService_Register(t *testing.T) {
	t.Parallel()
	type args struct {
		user *repository.RegisteredUser
	}
	tests := []struct {
		name      string
		args      args
		setup     func(repo *mock_repository.MockUserRepository, args args)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.args)
			uc := NewUserService(repo)
			// Assertion
			tt.assertion(t, uc.Register(tt.args.user), fmt.Sprintf("userService.Register(%v)", tt.args.user))
		})
	}
}
