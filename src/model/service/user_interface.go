package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainInterface{userRepository}
}

type userDomainInterface struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) // Recebe um objeto do usuário e retorna um erro se houver
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr                      // Recebe uma string com o ID do usuário e um objeto do usuário para atualizar
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)                     // Recebe uma string com o ID do usuário e retorna um objeto do usuário e um erro se houver
	DeleteUser(string) *rest_err.RestErr                                                 // Recebe uma string com o ID do usuário e retorna um erro se houver
}
