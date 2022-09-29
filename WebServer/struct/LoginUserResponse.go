package _struct

import (
	GrpcUserService "WebServer/models/UserService"
	"encoding/json"
)

type LoginUserResponse struct {
	UserId    int32 `json:"UserId"`
	ErrorCode int32 `json:"ErrorCode"`
}

func GetLoginUserResponse(response GrpcUserService.LoginUserResponse) LoginUserResponse {
	return LoginUserResponse{
		ErrorCode: response.ErrorCode,
		UserId:    response.UserId,
	}
}

func GetReturnStringLoginUserResponse(response GrpcUserService.LoginUserResponse) string {
	res := GetLoginUserResponse(response)
	returnString, _ := json.Marshal(res)
	return string(returnString)
}
