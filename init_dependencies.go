package main

import (
	"github.com/kayquesza/gocrud-auth-api/src/controller"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository"
	"github.com/kayquesza/gocrud-auth-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database, // Recebe uma conexão com o banco de dados MongoDB
) controller.UserControllerInterface {
	repository := repository.NewUserRepository(database)  // Cria uma instância do repositório de usuários, pss. a conexão com o banco
	service := service.NewUserDomainService(repository)   // Cria uma instância do serviço de domínio de usuários, pss. o repositório criado
	return controller.NewUserControllerInterface(service) // Cria uma instância do controlador de usuários, pss. o serviço configurado
}
