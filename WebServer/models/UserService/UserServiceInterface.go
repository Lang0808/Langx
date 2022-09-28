package GrpcUserService

type UserServiceInterface interface {
	LoginUser(request *LoginUserRequest) *LoginUserResponse
	RegisterUser(request *RegisterUserRequest) *RegisterUserResponse
}
