// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package BlogDB

import (
	"context"
	"database/sql"
)

const createBlog = `-- name: CreateBlog :execresult
INSERT INTO blogs ( authorId, title, content, active, timeCreate)
VALUES
(?, ?, ?, ?, ?)
`

type CreateBlogParams struct {
	Authorid   int32
	Title      string
	Content    string
	Active     interface{}
	Timecreate sql.NullInt64
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBlog,
		arg.Authorid,
		arg.Title,
		arg.Content,
		arg.Active,
		arg.Timecreate,
	)
}
