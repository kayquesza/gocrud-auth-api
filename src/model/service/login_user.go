package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// Função que faz login de um usuário
func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface, 
) (model.UserDomainInterface, string, *rest_err.RestErr) { // Retorna o domínio de usuário, o token e um erro

	logger.Info("Initiating loginUser", 
		zap.String("journey", "loginUser")) 

	user, err := ud.userRepository.FindUserByEmail(userDomain.GetEmail()) // Busca o usuário por email
	if err != nil {
		return nil, "", rest_err.NewForbiddenError("User or password invalid") // Retorna um erro de usuário ou senha inválidos
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(userDomain.GetPassword())) // Compara a senha do usuário com a senha do domínio
	if errBcrypt != nil {                                                                                    
		return nil, "", rest_err.NewForbiddenError("User or password invalid") 
	}

	token, err := user.GenerateToken() // Gera o token do usuário
	if err != nil {                    
		return nil, "", err // retorna um nulo; uma string vazia, por não ter token gerado; e o erro
	}

	logger.Info("loginUser service executed successfully", 
		zap.String("userId", user.GetID()), 
		zap.String("journey", "loginUser")) 
	return user, token, nil 
}
