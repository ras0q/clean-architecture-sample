package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/ras0q/clean-architecture-sample/2_interface/handler"
)

type context struct {
	echo.Context // handler.Context<interface>を満たす
}

func f(next func(handler.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(context{c})
	}
}
