package GrpcUserService

import (
	"WebServer/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
	Config      config.Config
}

var Instance ZUserServiceClient

func init() {
	Config := config.Instance
	target := fmt.Sprintf(
		"%v:%v",
		Config.GetConfig("UserServiceHost"),
		Config.GetConfig("UserServicePort"),
	)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Connect to UserService failed")
		log.Fatal(err)
	}
	innerClient := NewUserServiceClient(conn)
	Instance = ZUserServiceClient{
		innerClient: innerClient,
		Config:      Config,
	}
}

func (c ZUserServiceClient) LoginUser(request *LoginUserRequest) *LoginUserResponse {
	response, _ := c.innerClient.LoginUser(context.Background(), request)
	return response
}

func (c ZUserServiceClient) RegisterUser(request *RegisterUserRequest) *RegisterUserResponse {
	response, _ := c.innerClient.RegisterUser(context.Background(), request)
	fmt.Printf("Register %v\n", response)
	return response
}
