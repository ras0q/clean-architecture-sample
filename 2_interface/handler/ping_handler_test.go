package handler_test

import (
	"fmt"
	"testing"

	"github.com/Ras96/clean-architecture-sample/2_interface/handler"
	"github.com/Ras96/clean-architecture-sample/2_interface/handler/mock_handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_pingHandler_Ping(t *testing.T) {
	t.Parallel()
	type args struct {
		c *mock_handler.MockContext
	}
	tests := []struct {
		name      string
		args      args
		setup     func(args args)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{},
			setup: func(args args) {
				args.c.EXPECT().String(200, "pong!").Return(nil)
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
			tt.setup(tt.args)
			h := handler.NewPingHandler()
			// Assertion
			tt.assertion(t, h.Ping(tt.args.c), fmt.Sprintf("pingHandler.Ping(%v)", tt.args.c))
		})
	}
}
