package database_test //TODO: import cycleになるので移動する

import (
	"fmt"
	"testing"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"github.com/Ras96/clean-architecture-sample/2_interface/database/mock_database"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userRepository_FindAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		want      []*model.User
		setup     func(SQLHandler *mock_database.MockSQLHandler, want []*model.User)
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
			setup: func(sqlHandler *mock_database.MockSQLHandler, want []*model.User) {
				users := make([]*model.User, 0)
				sqlHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*users = want

					return sqlHandler
				})
				sqlHandler.EXPECT().Error().Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			want: nil,
			setup: func(sqlHandler *mock_database.MockSQLHandler, want []*model.User) {
				users := make([]*model.User, 0)
				sqlHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*users = want

					return sqlHandler
				})
				sqlHandler.EXPECT().Error().Return(random.Error())
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			sqlHandler := mock_database.NewMockSQLHandler(ctrl)
			tt.setup(sqlHandler, tt.want)
			ur := database.NewUserRepository(sqlHandler)
			// Assertion
			got, err := ur.FindAll()
			tt.assertion(t, err, fmt.Sprintf("userRepository.FindAll()"))
			assert.Equalf(t, tt.want, got, "userRepository.FindAll()")
		})
	}
}

func Test_userRepository_FindByID(t *testing.T) {
	t.Parallel()
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name      string
		args      args
		want      *model.User
		setup     func(sqlHandler *mock_database.MockSQLHandler, args args, want *model.User)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id: random.UUID(),
			},
			want: &model.User{
				ID:    uuid.Nil, // setupでargs.idと揃える
				Name:  random.AlphaNumeric(5),
				Email: random.Email(),
			},
			setup: func(sqlHandler *mock_database.MockSQLHandler, args args, want *model.User) {
				want.ID = args.id
				user := model.User{ID: args.id}
				sqlHandler.EXPECT().First(&user, gomock.Any()).DoAndReturn(func(user *model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*user = *want

					return sqlHandler
				})
				sqlHandler.EXPECT().Error().Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			args: args{
				id: random.UUID(),
			},
			want: nil,
			setup: func(sqlHandler *mock_database.MockSQLHandler, args args, want *model.User) {
				user := model.User{ID: args.id}
				sqlHandler.EXPECT().First(&user, gomock.Any()).Return(sqlHandler)
				sqlHandler.EXPECT().Error().Return(random.Error())
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			sqlHandler := mock_database.NewMockSQLHandler(ctrl)
			tt.setup(sqlHandler, tt.args, tt.want)
			ur := database.NewUserRepository(sqlHandler)
			// Assertion
			got, err := ur.FindByID(tt.args.id)
			tt.assertion(t, err, fmt.Sprintf("userRepository.FindByID(%v)", tt.args.id))
			assert.Equalf(t, tt.want, got, "userRepository.FindByID(%v)", tt.args.id)
		})
	}
}

func Test_userRepository_Register(t *testing.T) {
	t.Parallel()
	type args struct {
		user *repository.RegisteredUser
	}
	tests := []struct {
		name      string
		args      args
		setup     func(sqlHandler *mock_database.MockSQLHandler, args args)
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
			setup: func(sqlHandler *mock_database.MockSQLHandler, args args) {
				sqlHandler.EXPECT().Create(args.user).Return(sqlHandler)
				sqlHandler.EXPECT().Error().Return(nil)
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
			setup: func(sqlHandler *mock_database.MockSQLHandler, args args) {
				sqlHandler.EXPECT().Create(args.user).Return(sqlHandler)
				sqlHandler.EXPECT().Error().Return(random.Error())
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			sqlHandler := mock_database.NewMockSQLHandler(ctrl)
			tt.setup(sqlHandler, tt.args)
			ur := database.NewUserRepository(sqlHandler)
			// Assertion
			tt.assertion(t, ur.Register(tt.args.user), fmt.Sprintf("userRepository.Register(%v)", tt.args.user))
		})
	}
}
