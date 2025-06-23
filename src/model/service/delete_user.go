package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

// Função que deleta um usuário
func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr { // Retorna um erro
	logger.Info("Initiating deleteUser method in UserDomain", // Mensagem de log
		zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário

	err := ud.userRepository.DeleteUser(userId) // Deleta o usuário no banco de dados
	if err != nil {                             // Se houver algum erro, retorna um erro
		logger.Error("Error trying to call repository", err, // Mensagem de log
			zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário
		return err // Retorna um erro
	}

	logger.Info("deleteUser service executed successfully", // Mensagem de log
		zap.String("userId", userId),        // ID do usuário
		zap.String("journey", "deleteUser")) // Jornada da deleção de um usuário
	return nil // Retorna nil
}
