package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que atualiza um usuário
func (ud *userDomainService) UpdateUser(
	userId string, // ID do usuário
	userDomain model.UserDomainInterface, // Domínio de usuário
) *rest_err.RestErr { // Retorna um erro
	logger.Info("Initiating updateUser method in UserDomain", // Mensagem de log
		zap.String("journey", "updateUser")) // Jornada da atualização de um usuário

	err := ud.userRepository.UpdateUser(userId, userDomain) // Atualiza o usuário no banco de dados
	if err != nil {                                         // Se houver algum erro, retorna um erro
		logger.Error("Initiating updateUser method in UserDomain", err, // Mensagem de log
			zap.String("journey", "updateUser")) // Jornada da atualização de um usuário
		return err // Retorna um erro
	}

	logger.Info("updateUser service executed successfully", // Mensagem de log
		zap.String("userId", userId),        // ID do usuário
		zap.String("journey", "updateUser")) // Jornada da atualização de um usuário
	return nil // Retorna nil
}
