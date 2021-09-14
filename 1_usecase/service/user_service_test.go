package service_test

import (
	"testing"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository/mock_repository"
	"github.com/Ras96/clean-architecture-sample/1_usecase/service"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_userService_GetAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		want      []*model.User
		setup     func(repo *mock_repository.MockUserRepository, want []*model.User)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			want: []*model.User{
				{
					ID:    random.UUID(),
					Name:  random.AlphaNumeric(5),
					Email: random.Email(),
				},
			},
			setup: func(repo *mock_repository.MockUserRepository, want []*model.User) {
				repo.EXPECT().FindAll().Return(want, nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			want: nil,
			setup: func(repo *mock_repository.MockUserRepository, want []*model.User) {
				repo.EXPECT().FindAll().Return(want, gorm.ErrInvalidDB)
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.want)

			uc := service.NewUserService(repo)
			got, err := uc.GetAll()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
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
		{
			name: "success",
			args: args{
				id: random.UUID(),
			},
			want: &model.User{
				ID:    random.UUID(),
				Name:  random.AlphaNumeric(5),
				Email: random.Email(),
			},
			setup: func(repo *mock_repository.MockUserRepository, args args, want *model.User) {
				repo.EXPECT().FindByID(args.id).Return(want, nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			args: args{
				id: random.UUID(),
			},
			want: nil,
			setup: func(repo *mock_repository.MockUserRepository, args args, want *model.User) {
				repo.EXPECT().FindByID(args.id).Return(want, gorm.ErrInvalidDB)
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.args, tt.want)

			uc := service.NewUserService(repo)
			got, err := uc.GetByID(tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
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
		{
			name: "success",
			args: args{
				user: &repository.RegisteredUser{
					ID:    random.UUID(),
					Name:  random.AlphaNumeric(5),
					Email: random.Email(),
				},
			},
			setup: func(repo *mock_repository.MockUserRepository, args args) {
				repo.EXPECT().Register(args.user).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			args: args{
				user: &repository.RegisteredUser{
					ID:    random.UUID(),
					Name:  random.AlphaNumeric(5),
					Email: random.Email(),
				},
			},
			setup: func(repo *mock_repository.MockUserRepository, args args) {
				repo.EXPECT().Register(args.user).Return(gorm.ErrInvalidDB)
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repository.NewMockUserRepository(ctrl)
			tt.setup(repo, tt.args)

			uc := service.NewUserService(repo)

			tt.assertion(t, uc.Register(tt.args.user))
		})
	}
}
