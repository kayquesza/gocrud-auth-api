package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
)

// Função que cria um usuário
func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface, // Domínio de usuário
) (model.UserDomainInterface, *rest_err.RestErr) { // Retorna o domínio de usuário e um erro

	logger.Info("Initiating CreateUser method in UserDomain", // Mensagem de log
		zap.String("journey", "createUser")) // Jornada da criação de um usuário

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail()) // Busca o usuário por email
	if user != nil {                                             // Se o usuário for encontrado, retorna um erro
		return nil, rest_err.NewBadRequestError("Email is already registered in another account") // Retorna um erro de email já registrado em outra conta
	}

	userDomain.EncryptPassword()                                          // Criptografa a senha do usuário
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain) // Cria o usuário no banco de dados
	if err != nil {                                                       // Se houver algum erro, retorna um erro
		logger.Error("Initiating CreateUser method in UserDomain", err, // Mensagem de log
			zap.String("journey", "createUser")) // Jornada da criação de um usuário
		return nil, err // Retorna um erro
	}

	logger.Info("User created successfully in UserDomain", // Mensagem de log
		zap.String("userId", userDomainRepository.GetID()), // ID do usuário
		zap.String("journey", "createUser"))                // Jornada da criação de um usuário
	return userDomainRepository, nil // Retorna o domínio de usuário
}
