package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que cria um usuário
func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface, 
) (model.UserDomainInterface, *rest_err.RestErr) { 

	logger.Info("Initiating CreateUser method in UserDomain",
		zap.String("journey", "createUser")) 

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail()) // Busca o usuário por email
	if user != nil {                                            
		return nil, rest_err.NewBadRequestError("Email is already registered in another account") // Retorna um erro de email já registrado em outra conta
	}

	userDomain.EncryptPassword()                                          // Criptografa a senha do usuário
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain) // Cria o usuário no banco de dados / Repository
	if err != nil {                                                       
		logger.Error("Initiating CreateUser method in UserDomain", err,
			zap.String("journey", "createUser")) 
		return nil, err // Retorna um erro
	}

	logger.Info("User created successfully in UserDomain",
		zap.String("userId", userDomainRepository.GetID()), 
		zap.String("journey", "createUser"))                
	return userDomainRepository, nil 
}
