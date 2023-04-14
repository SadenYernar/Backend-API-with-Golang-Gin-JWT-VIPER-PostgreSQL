package main

import (
	"log"
	"project/config"
	"project/internal/controller"
	"project/internal/database"
	"project/internal/repository"
	"project/internal/server"
	"project/internal/service"
)

func main() {
	config, err := config.LoadConfig("/home/student/project/config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := database.SetupDB()
	if err != nil {
		log.Println(err)
	}

	// defer conn.Close()

	store := repository.NewRepository(conn)
	service := service.NewService(store)
	handler := controller.NewHandler(service)

	srv := new(server.Server)

	if err := srv.ServerRun(config.Host, config.Port, handler.Route()); err != nil {
		log.Fatalf("error server: %v", err)
	}
}
