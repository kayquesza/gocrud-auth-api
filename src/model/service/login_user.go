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
	userDomain model.UserDomainInterface, // Domínio de usuário
) (model.UserDomainInterface, string, *rest_err.RestErr) { // Retorna o domínio de usuário, o token e um erro

	logger.Info("Initiating loginUser", // Mensagem de log
		zap.String("journey", "loginUser")) // Jornada da autenticação de um usuário

	user, err := ud.userRepository.FindUserByEmail(userDomain.GetEmail()) // Busca o usuário por email
	if err != nil {
		return nil, "", rest_err.NewForbiddenError("User or password invalid") // Retorna um erro de usuário ou senha inválidos
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(userDomain.GetPassword())) // Compara a senha do usuário com a senha do domínio
	if errBcrypt != nil {                                                                                    // Se a senha do usuário não for igual à senha do domínio, retorna um erro
		return nil, "", rest_err.NewForbiddenError("User or password invalid") // Retorna um erro de usuário ou senha inválidos
	}

	token, err := user.GenerateToken() // Gera o token do usuário
	if err != nil {                    // Se houver algum erro, retorna um erro
		return nil, "", err // retorna um nulo; uma string vazia, por não ter token gerado; e o erro
	}

	logger.Info("loginUser service executed successfully", // Mensagem de log
		zap.String("userId", user.GetID()), // ID do usuário
		zap.String("journey", "loginUser")) // Jornada da autenticação de um usuário
	return user, token, nil // Retorna o domínio de usuário, o token e nil
}
