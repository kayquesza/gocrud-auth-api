package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que atualiza um usuário
func (ud *userDomainService) UpdateUser(
	userId string, 
	userDomain model.UserDomainInterface, // Domínio de usuário
) *rest_err.RestErr { 
	logger.Info("Initiating updateUser method in UserDomain", 
		zap.String("journey", "updateUser")) 

	err := ud.userRepository.UpdateUser(userId, userDomain) // Atualiza o usuário no banco de dados
	if err != nil {                                         
		logger.Error("Initiating updateUser method in UserDomain", err, 
			zap.String("journey", "updateUser")) 
		return err 
	}

	logger.Info("updateUser service executed successfully", 
		zap.String("userId", userId),        
		zap.String("journey", "updateUser")) 
	return nil 
}
