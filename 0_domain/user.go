package domain

import (
	"github.com/google/uuid"
)

type User struct {
	id    uuid.UUID
	name  string
	email string
}

func NewUser(id uuid.UUID, name string, email string) User {
	return User{id, name, email}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}
