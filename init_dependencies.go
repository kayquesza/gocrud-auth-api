package main

import (
	"github.com/kayquesza/gocrud-auth-api/src/controller"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository"
	"github.com/kayquesza/gocrud-auth-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	// Iniciar as dependências
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	return controller.NewUserControllerInterface(service)

}
