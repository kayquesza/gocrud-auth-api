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

func (uc userControllerInterface) UpdateUser(c *gin.Context) { // Função que recebe um contexto do Gin e atualiza um usuário
	logger.Info("Init UpdateUser Controller", // Mensagem de log
		zap.String("journey", "updateUser"), // Jornada da atualização de um usuário
	)
	var userRequest request.UserUpdateRequest // Variável para armazenar o corpo da requisição

	if err := c.ShouldBindJSON(&userRequest); err != nil { // Se houver algum erro, retornará um erro 400 com a mensagem de erro
		logger.Error("Error trying to validade user info", err, // Mensagem de log
			zap.String("journey", "updateUser")) // Jornada da atualização de um usuário
		errRest := validation.ValidadeUserError(err) // Valida o usuário

		c.JSON(errRest.Code, errRest) // Retorna o erro 400 com a mensagem de erro
		return
	}

	userId := c.Param("userId")                                  // Obtém o ID do usuário da requisição
	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value") // Cria um erro de requisição malformada
		c.JSON(errRest.Code, errRest)                                                 // Retorna o erro de requisição malformada
	}

	domain := model.NewUserUpdateDomain( // Cria um novo domínio de atualização de usuário
		userRequest.Name, // Nome do usuário
		userRequest.Age,  // Idade do usuário
	)

	err := uc.service.UpdateUser(userId, domain) // Atualiza o usuário
	if err != nil {                              // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call UpdateUser service.", err, // Mensagem de log
			zap.String("journey", "updateUser")) // Jornada da atualização de um usuário
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User created succesfully", // Mensagem de log
		zap.String("userId", userId),        // ID do usuário
		zap.String("journey", "updateUser")) // Jornada da atualização de um usuário

	c.Status(http.StatusOK) // Retorna o status 200 OK
}
