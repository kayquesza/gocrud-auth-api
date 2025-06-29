package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/model/service"
)

// Função que recebe um serviço de domínio de usuário e retorna um controlador de usuário
func NewUserControllerInterface( 
	serviceInterface service.UserDomainService, 
) UserControllerInterface { 
	return &userControllerInterface{ 
		service: serviceInterface, 
	}
}

// Interface que define os métodos que o controlador de usuário deve implementar
type UserControllerInterface interface {
	FindUserByID(c *gin.Context)    // Método que busca um usuário por ID
	FindUserByEmail(c *gin.Context) // Método que busca um usuário por email

	DeleteUser(c *gin.Context) // Método que deleta um usuário
	CreateUser(c *gin.Context) // Método que cria um usuário
	UpdateUser(c *gin.Context) // Método que atualiza um usuário
	LoginUser(c *gin.Context)  // Método que faz login de um usuário
}

// Struct que define o controlador de usuário
type userControllerInterface struct {
	service service.UserDomainService 
}
