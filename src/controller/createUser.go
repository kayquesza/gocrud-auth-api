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
	UserDomainInterface model.UserDomainInterface // Interface para o domínio de usuário
)

func (uc userControllerInterface) CreateUser(c *gin.Context) { // Função que recebe um contexto do Gin e cria um usuário
	logger.Info("Init CreatUser Controller", // Registra o início da criação de um usuário
		zap.String("journey", "createUser"), // Registra a jornada da criação de um usuário
	)
	var userRequest request.UserRequest // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { // Se houver algum erro, retornará um erro 400 com a mensagem de erro
		logger.Error("Error trying to validade user info", err, // Registra o erro de validação do usuário
			zap.String("journey", "createUser")) // Registra a jornada da criação de um usuário
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	domain := model.NewUserDomain( // Cria um novo domínio de usuário
		userRequest.Email,    // Email do usuário
		userRequest.Password, // Senha do usuário
		userRequest.Name,     // Nome do usuário
		userRequest.Age,      // Idade do usuário
	)

	domainResult, err := uc.service.CreateUserService(domain) // Cria um novo usuário
	if err != nil {                                           // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call CreatUser service.", err, // Registra o erro de criação do usuário
			zap.String("journey", "createUser")) // Registra a jornada da criação de um usuário
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", // Registra o sucesso da criação do usuário
		zap.String("userId", domainResult.GetID()), // Registra o ID do usuário
		zap.String("journey", "createUser"))        // Registra a jornada da criação de um usuário

	c.JSON(http.StatusOK, view.ConvertDomainToResponse( // Converte o domínio de usuário para uma resposta
		domainResult, // Domínio de usuário
	)) // Retorna o usuário criado com sucesso
}
