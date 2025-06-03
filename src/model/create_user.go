package model

import (
	"fmt"

	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Initiating CreateUser method in UserDomain", zap.String("journey", "createUser"))
	// Implementação da lógica de criação de usuário

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
