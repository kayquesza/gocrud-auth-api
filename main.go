package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/database/mongodb"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/controller/routes"
)

func main() {
	logger.Info("Starting the server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err.Error())
		// Caso não consiga conectar ao banco de dados, o log irá capturar o erro e o servidor não iniciará
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	// Caso tenha algo que impeça o servidor de iniciar, como a porta já em uso, o log irá capturar o erro

}
