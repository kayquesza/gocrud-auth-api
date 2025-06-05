package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/database/mongodb"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/controller"
	"github.com/kayquesza/gocrud-auth-api/src/controller/routes"
	"github.com/kayquesza/gocrud-auth-api/src/model/service"
)

func main() {
	logger.Info("Starting the server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()
	// Inicializa a conexão com o MongoDB

	// Iniciar as dependências
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	// Caso tenha algo que impeça o servidor de iniciar, como a porta já em uso, o log irá capturar o erro

}
