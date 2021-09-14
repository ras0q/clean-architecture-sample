package handler_test

import (
	"testing"

	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler/mock_handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewAPI(t *testing.T) {
	t.Parallel()
	type args struct {
		ping handler.PingHandler
		user handler.UserHandler
	}
	tests := []struct {
		name  string
		args  args
		want  handler.API
		setup func(args args, want *handler.API)
	}{
		{
			name:  "success",
			args:  args{},
			want:  handler.API{},
			setup: func(args args, want *handler.API) {
				want.Ping = args.ping
				want.User = args.user
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Setup mock
			ctrl := gomock.NewController(t)
			tt.args.ping = mock_handler.NewMockPingHandler(ctrl)
			tt.args.user = mock_handler.NewMockUserHandler(ctrl)
			tt.setup(tt.args, &tt.want)
			// Assertion
			assert.Equalf(t, tt.want, handler.NewAPI(tt.args.ping, tt.args.user), "NewAPI(%v, %v)", tt.args.ping, tt.args.user)
		})
	}
}
