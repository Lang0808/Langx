package GrpcBlogService

import (
	"BlogService/errors"
	"context"
)

type ZBlogServiceClient struct {
	innerClient BlogServiceClient
}

func (z *ZBlogServiceClient) CreatePost(request *CreatePostRequest) CreatePostResponse {
	res, _ := z.innerClient.CreatePost(context.Background(), request)
	if res == nil {
		return CreatePostResponse{
			ErrorCode: errors.CANNOT_CONNECT_TO_BLOG_SERVICE,
		}
	}
	return *res
}

func (z *ZBlogServiceClient) GetPost(request *GetPostRequest) GetPostResponse {
	res, _ := z.innerClient.GetPost(context.Background(), request)
	if res == nil {
		return GetPostResponse{
			ErrorCode: errors.CANNOT_CONNECT_TO_BLOG_SERVICE,
		}
	}
	return *res
}
