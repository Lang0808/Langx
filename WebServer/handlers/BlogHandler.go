package handlers

import (
	"WebServer/config"
	"WebServer/errors"
	"WebServer/messages"
	GrpcBlogService "WebServer/models/BlogService"
	_struct "WebServer/struct"
	"WebServer/utils"
	"fmt"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	Config           config.Config
	BlogServiceModel GrpcBlogService.BlogServiceInterface
}

func (h *BlogHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	identify := utils.AuthorizeUtilsInstance.GetCookieAuthorize(r)
	if identify.ErrorCode != int32(errors.SUCCESS) {
		message := messages.ApiMessage{
			ErrorCode: int(identify.ErrorCode),
			Message:   "",
			Data:      "",
		}
		fmt.Fprintf(w, messages.GetApiMessageString(message))
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	authorId := identify.UserId
	request := GrpcBlogService.CreatePostRequest{
		AuthorId: authorId,
		Title:    title,
		Content:  content,
	}

	res := h.BlogServiceModel.CreatePost(&request)
	fmt.Printf("%v %v\n", res.PostId, res.ErrorCode)
	returnString := _struct.GetReturnStringCreatePostResponse(res)
	fmt.Println(returnString)
	fmt.Fprintf(w, returnString)
}

func (h *BlogHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	identify := utils.AuthorizeUtilsInstance.GetCookieAuthorize(r)
	if identify.ErrorCode != int32(errors.SUCCESS) {
		message := messages.ApiMessage{
			ErrorCode: int(identify.ErrorCode),
			Message:   "",
			Data:      "",
		}
		fmt.Fprintf(w, messages.GetApiMessageString(message))
		return
	}
	postId := r.FormValue("PostId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		message := messages.ApiMessage{
			ErrorCode: int(errors.INCORRECT_PARAM),
			Message:   "",
			Data:      "",
		}
		fmt.Fprintf(w, messages.GetApiMessageString(message))
		return
	}
	request := GrpcBlogService.GetPostRequest{
		PostId: int32(id),
	}

	res := h.BlogServiceModel.GetPost(&request)
	fmt.Printf("%v \n", res.ErrorCode)
	returnString := _struct.GetReturnStringGetPostResponse(res)
	fmt.Println(returnString)
	fmt.Fprintf(w, returnString)
}
