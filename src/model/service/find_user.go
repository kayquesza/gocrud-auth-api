package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que busca um usuário por ID
func (ud *userDomainService) FindUserByIDServices(
	id string, 
) (model.UserDomainInterface, *rest_err.RestErr) { 
	logger.Info("Initiating findUserByID Services", 
		zap.String("journey", "findUserById")) 

	return ud.userRepository.FindUserByID(id) // Busca o usuário por ID
}

// Função que busca um usuário por email
func (ud *userDomainService) FindUserByEmailServices(
	email string, 
) (model.UserDomainInterface, *rest_err.RestErr) { 
	logger.Info("Initiating findUserByEmail Services", 
		zap.String("journey", "findUserById")) 

	return ud.userRepository.FindUserByEmail(email) // Busca o usuário por email
}
