package models

import (
	_struct "BlogService/struct"
	"context"
)

type UserDBModelInterface interface {
	CreatePost(context context.Context, request _struct.CreatePostParams) (int32, *_struct.CreatePostReturn)
}
