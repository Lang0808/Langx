// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package BlogDB

import (
	"database/sql"
)

type Blog struct {
	ID         sql.NullInt32
	Authorid   int32
	Title      string
	Content    string
	Active     interface{}
	Timecreate sql.NullInt64
}
