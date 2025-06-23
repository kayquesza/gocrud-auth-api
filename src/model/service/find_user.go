package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que busca um usuário por ID
func (ud *userDomainService) FindUserByIDServices(
	id string, // ID do usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating findUserByID Services", // Mensagem de log
		zap.String("journey", "findUserById")) // Jornada da busca de um usuário por ID

	return ud.userRepository.FindUserByID(id) // Busca o usuário por ID
}

// Função que busca um usuário por email
func (ud *userDomainService) FindUserByEmailServices(
	email string, // Email do usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro
	logger.Info("Initiating findUserByEmail Services", // Mensagem de log
		zap.String("journey", "findUserById")) // Jornada da busca de um usuário por email

	return ud.userRepository.FindUserByEmail(email) // Busca o usuário por email
}
