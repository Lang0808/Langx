package _struct

import (
	GrpcBlogService "WebServer/models/BlogService"
	"encoding/json"
)

type GetPostResponse struct {
	PostId     int32  `json:"PostId"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	AuthorId   int32  `json:"AuthorId"`
	AuthorName string `json:"AuthorName"`
}

func GetGetPostResponse(response GrpcBlogService.GetPostResponse) GetPostResponse {
	AuthorName := "Bao"
	return GetPostResponse{
		Title:      response.Title,
		Content:    response.Content,
		AuthorId:   response.AuthorId,
		AuthorName: AuthorName,
	}
}

func GetReturnStringGetPostResponse(response GrpcBlogService.GetPostResponse) string {
	res := GetGetPostResponse(response)
	returnString, _ := json.Marshal(res)
	return string(returnString)
}
