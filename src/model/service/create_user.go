package service

import (
	"fmt"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainInterface) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Initiating CreateUser method in UserDomain", zap.String("journey", "createUser"))
	// Implementação da lógica de criação de usuário

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
