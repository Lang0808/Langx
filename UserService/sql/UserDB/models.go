// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package UserDB

import (
	"database/sql"
)

type Password struct {
	Userid   sql.NullInt32
	Password string
}

type User struct {
	ID       sql.NullInt32
	Username string
}
