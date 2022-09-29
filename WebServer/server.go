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
	http.HandleFunc("/api/LoginUser", userHandler.LoginUser)
	dir, _ := os.Getwd()
	fmt.Println("current path :" + dir)
	// http.Handle("/Register/", http.FileServer(http.Dir("./FE/Register")))
	// http.Handle("/", http.StripPrefix("/Login/", http.FileServer(http.Dir("./FE/Login"))))
	// http.HandleFunc("/", pathHandler.InvalidPath)

	http.HandleFunc("/", serveFiles)

	port := Config.GetConfig("port")
	fmt.Printf("Create server at %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalf("Error when listen and serve\n")
		log.Fatal(err)
	}
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "./FE" + r.URL.Path
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
