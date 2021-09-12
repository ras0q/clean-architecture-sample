//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=./mock_$GOPACKAGE/mock_$GOFILE

package database

type SQLHandler interface {
	Find(out interface{}, where ...interface{}) SQLHandler
	First(out interface{}, where ...interface{}) SQLHandler
	Create(value interface{}) SQLHandler
	Error() error
}
