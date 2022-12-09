package repository

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	domain "github.com/ras0q/clean-architecture-sample/0_domain"
	"github.com/ras0q/clean-architecture-sample/1_usecase/repository"
	"github.com/ras0q/clean-architecture-sample/2_interface/database/mock_database"
	"github.com/ras0q/clean-architecture-sample/2_interface/repository/model"
	"github.com/ras0q/clean-architecture-sample/util/random"
	"github.com/stretchr/testify/assert"
)

// FindByIDなどで指定するID
var specificID = random.UUID()

func Test_userRepository_FindAll(t *testing.T) {
	t.Parallel()
	type fields struct {
		SQLHandler *mock_database.MockSQLHandler
	}
	tests := []struct {
		name      string
		fields    fields
		want      []*domain.User
		setup     func(f fields, want []*domain.User)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			want: []*domain.User{
				random.User(),
				random.User(),
				random.User(),
			},
			setup: func(f fields, want []*domain.User) {
				users := make([]*domain.User, 0)
				f.SQLHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*domain.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
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
			setup: func(f fields, want []*domain.User) {
				users := make([]*domain.User, 0)
				f.SQLHandler.EXPECT().Find(&users, gomock.Any()).DoAndReturn(func(users *[]*domain.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
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
		want      *domain.User
		setup     func(f fields, args args, want *domain.User)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id: specificID,
			},
			want: func() *domain.User {
				user := domain.NewUser(specificID, random.AlphaNumeric(5), random.Email())

				return &user
			}(), // TODO: もう少し簡潔に書きたい
			setup: func(f fields, args args, want *domain.User) {
				user := domain.NewUser(args.id, "", "")
				f.SQLHandler.EXPECT().First(&user, gomock.Any()).DoAndReturn(func(user *domain.User, any ...gomock.Matcher) *mock_database.MockSQLHandler {
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
				id: specificID,
			},
			want: nil,
			setup: func(f fields, args args, want *domain.User) {
				user := domain.NewUser(specificID, "", "")
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
				u := model.User{
					ID:    args.user.ID,
					Name:  args.user.Name,
					Email: args.user.Email,
				}
				f.SQLHandler.EXPECT().Create(&u).Return(f.SQLHandler)
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
				u := model.User{
					ID:    args.user.ID,
					Name:  args.user.Name,
					Email: args.user.Email,
				}
				f.SQLHandler.EXPECT().Create(&u).Return(f.SQLHandler)
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
