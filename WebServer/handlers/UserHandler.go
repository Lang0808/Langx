package handlers

import (
	"WebServer/config"
	GrpcUserService "WebServer/models/UserService"
	_struct "WebServer/struct"
	"fmt"
	"net/http"
)

type UserHandler struct {
	Config           config.Config
	UserServiceModel GrpcUserService.UserServiceInterface
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	request := GrpcUserService.RegisterUserRequest{
		Username: username,
		Password: password,
	}

	res := h.UserServiceModel.RegisterUser(&request)
	fmt.Printf("%v %v\n", res.UserId, res.ErrorCode)
	returnString := _struct.GetReturnString(*res)
	fmt.Println(returnString)
	fmt.Fprintf(w, returnString)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	request := GrpcUserService.LoginUserRequest{
		Username: username,
		Password: password,
	}

	res := h.UserServiceModel.LoginUser(&request)
	fmt.Printf("%v %v\n", res.UserId, res.ErrorCode)
	returnString := _struct.GetReturnStringLoginUserResponse(*res)
	fmt.Println(returnString)
	fmt.Fprintf(w, returnString)
}
