module github.com/Ras96/clean-architecture-sample

go 1.16 // TODO: github.com/google/wireがgo1.17に対応したら変更する

require (
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/google/wire v0.5.0
	github.com/labstack/echo/v4 v4.5.0
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)
