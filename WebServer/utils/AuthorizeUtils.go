package utils

import (
	"WebServer/config"
	"WebServer/errors"
	_struct "WebServer/struct"
	"fmt"
	"net/http"
	"time"
)

type AuthorizeUtils struct {
	c config.Config
}

var AuthorizeUtilsInstance *AuthorizeUtils

func init() {
	AuthorizeUtilsInstance = &AuthorizeUtils{
		c: config.Instance,
	}
}

func (a *AuthorizeUtils) AddCookieAuthorize(r *http.Request, i InfoInJwt) {
	token, err := Instance.GenerateToken(i)
	if err != nil {
		c := &http.Cookie{
			Path:    "/",
			Name:    "token2",
			Value:   token,
			Domain:  fmt.Sprintf("%v", a.c.GetConfig("host")),
			Expires: time.Now().Add(3400 * time.Minute),
			/*HttpOnly: true,*/
		}
		r.AddCookie(c)
		fmt.Println(token)
	}
}

func (a *AuthorizeUtils) GetCookieAuthorize(r *http.Request) _struct.Identify {
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
	info, err := Instance.DecodeToken(res.Value)
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
