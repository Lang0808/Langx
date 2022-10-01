package _struct

import (
	GrpcBlogService "WebServer/models/BlogService"
	"encoding/json"
)

type CreatePostResponse struct {
	PostId    int32 `json:"PostId"`
	ErrorCode int32 `json:"ErrorCode"`
}

func GetCreatePostResponse(response GrpcBlogService.CreatePostResponse) CreatePostResponse {
	return CreatePostResponse{
		ErrorCode: response.ErrorCode,
		PostId:    response.PostId,
	}
}

func GetReturnStringCreatePostResponse(response GrpcBlogService.CreatePostResponse) string {
	res := GetCreatePostResponse(response)
	returnString, _ := json.Marshal(res)
	return string(returnString)
}
