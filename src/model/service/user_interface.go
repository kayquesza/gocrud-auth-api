package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository"
)

// Função que cria um novo serviço de domínio de usuário
func NewUserDomainService(
	userRepository repository.UserRepository, // Repositório de usuários
) UserDomainService { // Retorna um serviço de domínio de usuário
	return &userDomainService{userRepository: userRepository} // Retorna um serviço de domínio de usuário
}

// Struct que define o serviço de domínio de usuário
type userDomainService struct {
	userRepository repository.UserRepository // Repositório de usuários
}

// Interface que define os métodos que o serviço de domínio de usuário deve implementar
type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) ( // Método que cria um usuário
		model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro

	FindUserByIDServices( // Método que busca um usuário por ID
		id string, // ID do usuário
	) (model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro
	FindUserByEmailServices( // Método que busca um usuário por email
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr) // Retorna o domínio de usuário e um erro

	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr // Recebe uma string com o ID do usuário e um objeto do usuário para atualizar
	DeleteUser(string) *rest_err.RestErr                            // Retorna um erro

	LoginUserService( // Método que faz login de um usuário
		userDomain model.UserDomainInterface, // Domínio de usuário
	) (model.UserDomainInterface, string, *rest_err.RestErr) // Retorna o domínio de usuário, o token e um erro
}
