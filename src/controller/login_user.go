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

// Função que recebe um contexto do Gin e faz login de um usuário
func (uc *userControllerInterface) LoginUser(c *gin.Context) { 
	logger.Info("Init loginUser Controller", 
		zap.String("journey", "loginUser"), 
	)

	var userRequest request.UserLogin // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { 
		logger.Error("Error trying to validade user info", err, 
			zap.String("journey", "loginUser")) 
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	domain := model.NewUserLoginDomain( // Cria um novo domínio de login de usuário
		userRequest.Email,    
		userRequest.Password, 
	)

	domainResult, token, err := uc.service.LoginUserService(domain) // Faz o login do usuário
	if err != nil {                                               
		logger.Error("Error trying to call loginUser service.", err, 
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", 
		zap.String("userId", domainResult.GetID()), 
		zap.String("journey", "loginUser"))        

	c.Header("Authorization", token) // Define o token de autenticação

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult, // Domínio de usuário
	))

}
