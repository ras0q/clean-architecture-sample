//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package handler

type Context interface {
	Param(name string) string
	Bind(i interface{}) error
	Validate(i interface{}) error
	NoContent(code int) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
}
