package GrpcUserService

import (
	"UserService/errors"
	"UserService/sql/models"
	_struct "UserService/struct"
	"context"
	"fmt"
)

type ZUserServiceServer struct {
	UserDBModel models.UserDBModelInterface
}

func (s ZUserServiceServer) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	return &LoginUserResponse{
		UserId:    -1,
		ErrorCode: errors.UNSUPPORTED,
	}, nil
}
func (s ZUserServiceServer) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	fmt.Println(request)
	params := _struct.RegisterUserParams{
		Username: request.Username,
		Password: request.Password,
	}
	ErrorCode, UserID, _ := s.UserDBModel.RegisterUser(ctx, params)
	fmt.Printf("%v %v\n", ErrorCode, UserID)
	res := RegisterUserResponse{
		UserId:    UserID,
		ErrorCode: ErrorCode,
	}
	fmt.Printf("%v\n", res)
	return &res, nil
}

func (s ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {

}
