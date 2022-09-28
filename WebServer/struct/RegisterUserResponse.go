package _struct

import (
	GrpcUserService "WebServer/models/UserService"
	"encoding/json"
)

type RegisterUserResponse struct {
	ErrorCode int32 `json:"ErrorCode"`
}

func GetRegisterUserResponse(response GrpcUserService.RegisterUserResponse) RegisterUserResponse {
	return RegisterUserResponse{
		ErrorCode: response.ErrorCode,
	}
}

func GetReturnString(response GrpcUserService.RegisterUserResponse) string {
	res := GetRegisterUserResponse(response)
	returnString, _ := json.Marshal(res)
	return string(returnString)
}
