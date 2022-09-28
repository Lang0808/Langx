package main

import (
	"WebServer/config"
	"WebServer/handlers"
	GrpcUserService "WebServer/models/UserService"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	Config := config.Instance
	userHandler := handlers.UserHandler{Config: Config, UserServiceModel: GrpcUserService.Instance}
	// pathHandler := handlers.PathHandler{}

	http.HandleFunc("/api/RegisterUser", userHandler.RegisterUser)
	dir, _ := os.Getwd()
	fmt.Println("current path :" + dir)
	http.Handle("/", http.StripPrefix("/Register/", http.FileServer(http.Dir("./FE/Register"))))
	// http.HandleFunc("/", pathHandler.InvalidPath)

	port := Config.GetConfig("port")
	fmt.Printf("Create server at %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalf("Error when listen and serve\n")
		log.Fatal(err)
	}
}
