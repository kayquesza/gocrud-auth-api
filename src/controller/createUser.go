package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/validation"
	"github.com/kayquesza/gocrud-auth-api/src/controller/model/request"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/view"
	"go.uber.org/zap"
)

var ( // Variáveis de ambiente para o controlador de usuário
	UserDomainInterface model.UserDomainInterface 
)

func (uc userControllerInterface) CreateUser(c *gin.Context) { // Função que recebe um contexto do Gin e cria um usuário
	logger.Info("Init CreatUser Controller", 
		zap.String("journey", "createUser"), 
	)
	var userRequest request.UserRequest // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { 
		logger.Error("Error trying to validade user info", err, 
			zap.String("journey", "createUser")) 
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	domain := model.NewUserDomain( // Cria um novo domínio de usuário
		userRequest.Email,    
		userRequest.Password, 
		userRequest.Name,     
		userRequest.Age,      
	)

	domainResult, err := uc.service.CreateUserService(domain) // Cria um novo usuário
	if err != nil {                                          
		logger.Error("Error trying to call CreatUser service.", err, 
			zap.String("journey", "createUser")) 
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", // Registra o sucesso da criação do usuário
		zap.String("userId", domainResult.GetID()), // Registra o ID do usuário
		zap.String("journey", "createUser"))       

	c.JSON(http.StatusOK, view.ConvertDomainToResponse( // Converte o domínio de usuário para uma resposta
		domainResult, // Domínio de usuário
	)) // Retorna o usuário criado com sucesso
}
