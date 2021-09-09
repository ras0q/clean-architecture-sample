package database

type SQLHandler interface {
	Find(out interface{}, where ...interface{}) SQLHandler
	First(out interface{}, where ...interface{}) SQLHandler
	Create(value interface{}) SQLHandler
	Error() error
}
