package models

import (
	"platzi/chat/app/lib"
	"time"
)

type User struct {
	Id        int64     `db:"id,omitempty"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	*lib.Access
}
