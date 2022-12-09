package alt

import (
	"github.com/gin-gonic/gin"
	"github.com/ras0q/clean-architecture-sample/2_interface/handler"
)

type context struct {
	*gin.Context
}

func f(next func(handler.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		next(context{c})
	}
}

func (c context) Validate(i interface{}) error {
	return c.Context.ShouldBind(i) //TODO: 試してない
}

func (c context) NoContent(code int) error {
	c.Context.Status(code)

	return nil
}

func (c context) String(code int, s string) error {
	c.Context.String(code, s)

	return nil
}

func (c context) JSON(code int, i interface{}) error {
	c.Context.JSON(code, i)

	return nil
}
