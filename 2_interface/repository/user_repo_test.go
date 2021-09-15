package repository

import (
	"fmt"
	"testing"

	"github.com/Ras96/clean-architecture-sample/0_domain/model"
	"github.com/Ras96/clean-architecture-sample/0_domain/repository"
	"github.com/Ras96/clean-architecture-sample/2_interface/database/mock_database"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userRepository_FindAll(t *testing.T) {
	t.Parallel()
	type fields struct {
		SQLHandler *mock_database.MockSQLHandler
	}
	tests := []struct {
		name      string
		fields    fields
		want      []*model.User
		setup     func(f fields, want []*model.User)
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
			setup: func(f fields, want []*model.User) {
				users := make([]*model.User, 0)
				f.SQLHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*users = want

					return f.SQLHandler
				})
				f.SQLHandler.EXPECT().Error().Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			want: nil,
			setup: func(f fields, want []*model.User) {
				users := make([]*model.User, 0)
				f.SQLHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*users = want

					return f.SQLHandler
				})
				f.SQLHandler.EXPECT().Error().Return(random.Error())
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
			tt.fields = fields{
				SQLHandler: mock_database.NewMockSQLHandler(ctrl),
			}
			tt.setup(tt.fields, tt.want)
			ur := NewUserRepository(tt.fields.SQLHandler)
			// Assertion
			got, err := ur.FindAll()
			tt.assertion(t, err, fmt.Sprintf("userRepository.FindAll()"))
			assert.Equalf(t, tt.want, got, "userRepository.FindAll()")
		})
	}
}

func Test_userRepository_FindByID(t *testing.T) {
	t.Parallel()
	type fields struct {
		SQLHandler *mock_database.MockSQLHandler
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *model.User
		setup     func(f fields, args args, want *model.User)
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
			setup: func(f fields, args args, want *model.User) {
				want.ID = args.id
				user := model.User{ID: args.id}
				f.SQLHandler.EXPECT().First(&user, gomock.Any()).DoAndReturn(func(user *model.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
					*user = *want

					return f.SQLHandler
				})
				f.SQLHandler.EXPECT().Error().Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			args: args{
				id: random.UUID(),
			},
			want: nil,
			setup: func(f fields, args args, want *model.User) {
				user := model.User{ID: args.id}
				f.SQLHandler.EXPECT().First(&user, gomock.Any()).Return(f.SQLHandler)
				f.SQLHandler.EXPECT().Error().Return(random.Error())
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
			tt.fields = fields{
				SQLHandler: mock_database.NewMockSQLHandler(ctrl),
			}
			tt.setup(tt.fields, tt.args, tt.want)
			ur := NewUserRepository(tt.fields.SQLHandler)
			// Assertion
			got, err := ur.FindByID(tt.args.id)
			tt.assertion(t, err, fmt.Sprintf("userRepository.FindByID(%v)", tt.args.id))
			assert.Equalf(t, tt.want, got, "userRepository.FindByID(%v)", tt.args.id)
		})
	}
}

func Test_userRepository_Register(t *testing.T) {
	t.Parallel()
	type fields struct {
		SQLHandler *mock_database.MockSQLHandler
	}
	type args struct {
		user *repository.RegisteredUser
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f fields, args args)
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
			setup: func(f fields, args args) {
				f.SQLHandler.EXPECT().Create(args.user).Return(f.SQLHandler)
				f.SQLHandler.EXPECT().Error().Return(nil)
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
			setup: func(f fields, args args) {
				f.SQLHandler.EXPECT().Create(args.user).Return(f.SQLHandler)
				f.SQLHandler.EXPECT().Error().Return(random.Error())
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
			tt.fields = fields{
				SQLHandler: mock_database.NewMockSQLHandler(ctrl),
			}
			tt.setup(tt.fields, tt.args)
			ur := NewUserRepository(tt.fields.SQLHandler)
			// Assertion
			tt.assertion(t, ur.Register(tt.args.user), fmt.Sprintf("userRepository.Register(%v)", tt.args.user))
		})
	}
}
