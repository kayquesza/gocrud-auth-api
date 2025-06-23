package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/validation"
	"github.com/kayquesza/gocrud-auth-api/src/controller/model/request"
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que recebe um contexto do Gin e atualiza um usuário
func (uc userControllerInterface) UpdateUser(c *gin.Context) { 
	logger.Info("Init UpdateUser Controller", 
		zap.String("journey", "updateUser"), 
	)
	var userRequest request.UserUpdateRequest // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { 
		logger.Error("Error trying to validade user info", err, 
			zap.String("journey", "updateUser")) 
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	userId := c.Param("userId")                                  // Obtém o ID do usuário da requisição
	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value") 
		c.JSON(errRest.Code, errRest)                                                 // Retorna o erro de requisição malformada
	}

	domain := model.NewUserUpdateDomain( // Cria um novo domínio de atualização de usuário
		userRequest.Name, 
		userRequest.Age,  
	)

	err := uc.service.UpdateUser(userId, domain) // Atualiza o usuário
	if err != nil {                              
		logger.Error("Error trying to call UpdateUser service.", err, 
			zap.String("journey", "updateUser")) 
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", 
		zap.String("userId", userId),        
		zap.String("journey", "updateUser")) 

	c.Status(http.StatusOK) // Retorna o status 200 OK
}
