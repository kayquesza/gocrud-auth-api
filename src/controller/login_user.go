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

func (uc *userControllerInterface) LoginUser(c *gin.Context) { // Função que recebe um contexto do Gin e faz login de um usuário
	logger.Info("Init loginUser Controller", // Mensagem de log
		zap.String("journey", "loginUser"), // Jornada da autenticação de um usuário
	)

	var userRequest request.UserLogin // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { // Se houver algum erro, retornará um erro 400 com a mensagem de erro
		logger.Error("Error trying to validade user info", err, // Mensagem de log
			zap.String("journey", "loginUser")) // Jornada da autenticação de um usuário
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	domain := model.NewUserLoginDomain( // Cria um novo domínio de login de usuário
		userRequest.Email,    // Email do usuário
		userRequest.Password, // Senha do usuário
	)

	domainResult, token, err := uc.service.LoginUserService(domain) // Faz o login do usuário
	if err != nil {                                                 // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call loginUser service.", err, // Mensagem de log
			zap.String("journey", "loginUser")) // Jornada da autenticação de um usuário
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", // Mensagem de log
		zap.String("userId", domainResult.GetID()), // ID do usuário
		zap.String("journey", "loginUser"))         // Jornada da autenticação de um usuário

	c.Header("Authorization", token) // Define o token de autenticação

	c.JSON(http.StatusOK, view.ConvertDomainToResponse( // Converte o domínio de usuário para uma resposta
		domainResult, // Domínio de usuário
	))

}
