package handlers

import (
	"WebServer/config"
	"WebServer/errors"
	GrpcUserService "WebServer/models/UserService"
	_struct "WebServer/struct"
	"WebServer/utils"
	"fmt"
	"net/http"
	"time"
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
	cookie, _ := r.Cookie("token2")
	fmt.Println(cookie)
	cookie2, _ := r.Cookie("token")
	fmt.Println(cookie2)
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	request := GrpcUserService.LoginUserRequest{
		Username: username,
		Password: password,
	}

	res := h.UserServiceModel.LoginUser(&request)
	fmt.Printf("%v %v\n", res.UserId, res.ErrorCode)
	returnString := _struct.GetReturnStringLoginUserResponse(*res)
	cookie3 := h.AddCookieAuthorize(&w, utils.InfoInJwt{UserId: int(res.UserId)})
	http.SetCookie(w, cookie3)
	fmt.Println(returnString)
	fmt.Fprintf(w, returnString)
}

func (h *UserHandler) AddCookieAuthorize(w *http.ResponseWriter, i utils.InfoInJwt) *http.Cookie {
	token, err := utils.Instance.GenerateToken(i)
	if err != nil {
		return nil
	}
	c := &http.Cookie{
		Path:    "/",
		Name:    "token2",
		Value:   token,
		Domain:  fmt.Sprintf("%v", h.Config.GetConfig("host")),
		Expires: time.Now().Add(3400 * time.Minute),
	}
	return c

}

func (h *UserHandler) GetCookieAuthorize(r *http.Request) _struct.Identify {
	res, err := r.Cookie("token2")
	if err != nil {
		return _struct.Identify{
			ErrorCode: errors.CANNOT_GET_AUTHORIZE_COOKIE,
			UserId:    -1,
		}
	}
	if res == nil {
		return _struct.Identify{
			ErrorCode: errors.CANNOT_GET_AUTHORIZE_COOKIE,
			UserId:    -1,
		}
	}
	info, err := utils.Instance.DecodeToken(res.Value)
	if err != nil {
		return _struct.Identify{
			ErrorCode: errors.CANNOT_GET_AUTHORIZE_COOKIE,
			UserId:    -1,
		}
	}
	if info == nil {
		return _struct.Identify{
			ErrorCode: errors.CANNOT_GET_AUTHORIZE_COOKIE,
			UserId:    -1,
		}
	}
	return _struct.Identify{
		ErrorCode: int32(errors.SUCCESS),
		UserId:    int32(info.UserId),
	}
}
