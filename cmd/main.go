package main

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/config"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/controller"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/database"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/server"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/service"
	"log"
)

func main() {
	config, err := config.LoadConfig("/home/student/Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/config")
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
