package models

import (
	_struct "UserService/struct"
	"context"
)

type UserDBModelInterface interface {
	RegisterUser(ctx context.Context, request _struct.RegisterUserParams) (int32, int32, error)
}
