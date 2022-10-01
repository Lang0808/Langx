package GrpcBlogService

type BlogServiceInterface interface {
	CreatePost(request *CreatePostRequest) CreatePostResponse
	GetPost(request *GetPostRequest) GetPostResponse
}
