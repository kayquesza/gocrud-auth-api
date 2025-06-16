package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {

	logger.Info("Initiating loginUser",
		zap.String("journey", "loginUser"))

	user, err := ud.userRepository.FindUserByEmail(userDomain.GetEmail())
	if err != nil {
		return nil, "", rest_err.NewForbiddenError("User or password invalid")
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(userDomain.GetPassword()))
	if errBcrypt != nil {
		return nil, "", rest_err.NewForbiddenError("User or password invalid")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err // retorna um nulo; uma string vazia, por não ter token gerado; e o erro
	}

	logger.Info("loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))
	return user, token, nil
}
