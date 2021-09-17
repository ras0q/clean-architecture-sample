package handler_test

import (
	"fmt"
	"net/http"
	"testing"

	domain "github.com/Ras96/clean-architecture-sample/0_domain"
	"github.com/Ras96/clean-architecture-sample/1_usecase/service/mock_service"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler/mock_handler"
	"github.com/Ras96/clean-architecture-sample/util/random"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_userHandler_GetAll(t *testing.T) {
	t.Parallel()
	type fields struct {
		srv *mock_service.MockUserService
	}
	type args struct {
		c *mock_handler.MockContext
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
			args: args{},
			setup: func(f fields, args args) {
				users := []*domain.User{
					{
						ID:    random.UUID(),
						Name:  random.AlphaNumeric(5),
						Email: random.Email(),
					},
				}
				f.srv.EXPECT().GetAll().Return(users, nil)
				res := make([]*handler.UserRes, 0, len(users))
				for _, v := range users {
					res = append(res, &handler.UserRes{
						ID:   v.ID,
						Name: v.Name,
					})
				}
				args.c.EXPECT().JSON(http.StatusOK, res).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_500",
			args: args{},
			setup: func(f fields, args args) {
				err := random.Error()
				f.srv.EXPECT().GetAll().Return(nil, err)
				args.c.EXPECT().JSON(http.StatusInternalServerError, err.Error()).Return(nil)
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			tt.args.c = mock_handler.NewMockContext(ctrl)
			tt.fields = fields{
				srv: mock_service.NewMockUserService(ctrl),
			}
			tt.setup(tt.fields, tt.args)
			h := handler.NewUserHandler(tt.fields.srv)
			// Assertion
			tt.assertion(t, h.GetAll(tt.args.c), fmt.Sprintf("userHandler.GetAll(%v)", tt.args.c))
		})
	}
}

func Test_userHandler_GetByID(t *testing.T) {
	t.Parallel()
	type fields struct {
		srv *mock_service.MockUserService
	}
	type args struct {
		c *mock_handler.MockContext
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
			args: args{},
			setup: func(f fields, args args) {
				id := random.UUID()
				args.c.EXPECT().Param("id").Return(id.String())
				user := &domain.User{
					ID:    id,
					Name:  random.AlphaNumeric(5),
					Email: random.Email(),
				}
				f.srv.EXPECT().GetByID(id).Return(user, nil)
				args.c.EXPECT().JSON(http.StatusOK, &handler.UserDetailRes{
					UserRes: handler.UserRes{
						ID:   user.ID,
						Name: user.Name,
					},
					Email: user.Email,
				}).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_400",
			args: args{},
			setup: func(f fields, args args) {
				idstr := random.AlphaNumeric(32)
				args.c.EXPECT().Param("id").Return(idstr)
				_, err := uuid.FromString(idstr)
				args.c.EXPECT().JSON(http.StatusBadRequest, err.Error()).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_404",
			args: args{},
			setup: func(f fields, args args) {
				id := random.UUID()
				args.c.EXPECT().Param("id").Return(id.String())
				f.srv.EXPECT().GetByID(id).Return(nil, gorm.ErrRecordNotFound)
				args.c.EXPECT().NoContent(http.StatusNotFound).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_500",
			args: args{},
			setup: func(f fields, args args) {
				id := random.UUID()
				args.c.EXPECT().Param("id").Return(id.String())
				err := random.Error()
				f.srv.EXPECT().GetByID(id).Return(nil, err)
				args.c.EXPECT().JSON(http.StatusInternalServerError, err.Error()).Return(nil)
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			tt.args.c = mock_handler.NewMockContext(ctrl)
			tt.fields = fields{
				srv: mock_service.NewMockUserService(ctrl),
			}
			tt.setup(tt.fields, tt.args)
			h := handler.NewUserHandler(tt.fields.srv)
			// Assertion
			tt.assertion(t, h.GetByID(tt.args.c), fmt.Sprintf("userHandler.GetByID(%v)", tt.args.c))
		})
	}
}

func Test_userHandler_Register(t *testing.T) {
	t.Parallel()
	type fields struct {
		srv *mock_service.MockUserService
	}
	type args struct {
		c *mock_handler.MockContext
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
			args: args{},
			setup: func(f fields, args args) {
				req := handler.RegisterReq{}
				args.c.EXPECT().Bind(&req).Return(nil)
				f.srv.EXPECT().Register(gomock.Any()).Return(nil)
				args.c.EXPECT().NoContent(http.StatusCreated).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_400",
			args: args{},
			setup: func(f fields, args args) {
				req := handler.RegisterReq{}
				err := random.Error()
				args.c.EXPECT().Bind(&req).Return(err)
				args.c.EXPECT().JSON(http.StatusBadRequest, err.Error()).Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "error_500",
			args: args{},
			setup: func(f fields, args args) {
				req := handler.RegisterReq{}
				args.c.EXPECT().Bind(&req).Return(nil)
				err := random.Error()
				f.srv.EXPECT().Register(gomock.Any()).Return(err)
				args.c.EXPECT().JSON(http.StatusInternalServerError, err.Error()).Return(nil)
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			tt.args.c = mock_handler.NewMockContext(ctrl)
			tt.fields = fields{
				srv: mock_service.NewMockUserService(ctrl),
			}
			tt.setup(tt.fields, tt.args)
			h := handler.NewUserHandler(tt.fields.srv)
			// Assertion
			tt.assertion(t, h.Register(tt.args.c), fmt.Sprintf("userHandler.Register(%v)", tt.args.c))
		})
	}
}
