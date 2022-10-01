package GrpcBlogService

import (
	"BlogService/errors"
	"BlogService/sql/models"
	_struct "BlogService/struct"
	"context"
)

type ZBlogServiceServer struct {
	b *models.BlogDBModel
}

func (z ZBlogServiceServer) CreatePost(ctx context.Context, request *CreatePostRequest) (*CreatePostResponse, error) {
	params := _struct.CreatePostParams{
		AuthorId: request.AuthorId,
		Title:    request.Title,
		Content:  request.Content,
	}
	ErrorCode, res, err := z.b.CreatePost(ctx, params)
	if err != nil {
		return &CreatePostResponse{
			ErrorCode: ErrorCode,
			PostId:    -1,
		}, nil
	}
	return &CreatePostResponse{
		ErrorCode: errors.SUCCESS,
		PostId:    res.PostID,
	}, nil
}

func (z ZBlogServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return &GetPostResponse{
		ErrorCode: errors.ERROR_NOT_SUPPORTED_YET,
		AuthorId:  -1,
		Content:   "",
		Title:     "",
	}, nil
}

func (ZBlogServiceServer) mustEmbedUnimplementedUserServiceServer() {

}
