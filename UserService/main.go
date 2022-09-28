package main

import (
	"UserService/config"
	GrpcUserService "UserService/grpc"
	"UserService/sql/models"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	Config := config.Instance
	host := fmt.Sprintf(
		"%v:%v",
		Config.GetConfig("host"),
		Config.GetConfig("port"),
	)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", host)
	}

	s := grpc.NewServer()
	UserServiceServer := GrpcUserService.ZUserServiceServer{
		UserDBModel: models.Instance,
	}
	GrpcUserService.RegisterUserServiceServer(s, UserServiceServer)

	fmt.Printf("ZUserServiceServer is listening at %v", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
