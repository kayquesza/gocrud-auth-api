package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

// Função que deleta um usuário
func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr { 
	logger.Info("Initiating deleteUser method in UserDomain", 
		zap.String("journey", "deleteUser")) 

	err := ud.userRepository.DeleteUser(userId) // Deleta o usuário no banco de dados
	if err != nil {                             
		logger.Error("Error trying to call repository", err, 
			zap.String("journey", "deleteUser")) 
		return err // Retorna um erro
	}

	logger.Info("deleteUser service executed successfully", 
		zap.String("userId", userId),       
		zap.String("journey", "deleteUser")) 
	return nil 
}
