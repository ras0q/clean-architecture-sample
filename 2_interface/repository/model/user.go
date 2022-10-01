package model

import (
	"time"

	"github.com/google/uuid"
)

// usersテーブルの定義
type User struct {
	ID        uuid.UUID `gorm:"primary key"`
	Name      string    `gorm:"not null"`
	Email     string
	CreatedAt time.Time
}
