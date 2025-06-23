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

	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conecta ao banco de dados MongoDB
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err.Error())
		// Caso não consiga conectar ao banco de dados, o log irá capturar o erro e o servidor não iniciará
		return
	}

	// Inicializa as dependências
	userController := initDependencies(database)

	// Inicializa o servidor Gin
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	// Inicia o servidor na porta 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	// Caso tenha algo que impeça o servidor de iniciar, como a porta já em uso, o log irá capturar o erro

}
