package service

import (
	"fmt"
	"testing"

	domain "github.com/Ras96/clean-architecture-sample/0_domain"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository"
	"github.com/Ras96/clean-architecture-sample/1_usecase/repository/mock_repository"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_userService_GetAll(t *testing.T) {
	t.Parallel()
	type fields struct {
		user *mock_repository.MockUserRepository
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
				f.user.EXPECT().FindAll().Return(want, nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			want: nil,
			setup: func(f fields, want []*domain.User) {
				f.user.EXPECT().FindAll().Return(want, gorm.ErrInvalidDB)
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
				user: mock_repository.NewMockUserRepository(ctrl),
			}
			tt.setup(tt.fields, tt.want)
			s := NewUserService(tt.fields.user)
			// Assertion
			got, err := s.GetAll()
			tt.assertion(t, err, fmt.Sprintf("userService.GetAll()"))
			assert.Equalf(t, tt.want, got, "userService.GetAll()")
		})
	}
}

func Test_userService_GetByID(t *testing.T) {
	t.Parallel()
	type fields struct {
		user *mock_repository.MockUserRepository
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
				id: random.UUID(),
			},
			want: random.User(),
			setup: func(f fields, args args, want *domain.User) {
				f.user.EXPECT().FindByID(args.id).Return(want, nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "dbError",
			args: args{
				id: random.UUID(),
			},
			want: nil,
			setup: func(f fields, args args, want *domain.User) {
				f.user.EXPECT().FindByID(args.id).Return(want, gorm.ErrInvalidDB)
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
				user: mock_repository.NewMockUserRepository(ctrl),
			}
			tt.setup(tt.fields, tt.args, tt.want)
			s := NewUserService(tt.fields.user)
			// Assertion
			got, err := s.GetByID(tt.args.id)
			tt.assertion(t, err, fmt.Sprintf("userService.GetByID(%v)", tt.args.id))
			assert.Equalf(t, tt.want, got, "userService.GetByID(%v)", tt.args.id)
		})
	}
}

func Test_userService_Register(t *testing.T) {
	t.Parallel()
	type fields struct {
		user *mock_repository.MockUserRepository
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
				f.user.EXPECT().Register(args.user).Return(nil)
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
				f.user.EXPECT().Register(args.user).Return(gorm.ErrInvalidDB)
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
				user: mock_repository.NewMockUserRepository(ctrl),
			}
			tt.setup(tt.fields, tt.args)
			s := NewUserService(tt.fields.user)
			// Assertion
			tt.assertion(t, s.Register(tt.args.user), fmt.Sprintf("userService.Register(%v)", tt.args.user))
		})
	}
}
