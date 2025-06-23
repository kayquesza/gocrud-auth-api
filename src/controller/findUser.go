package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que recebe um contexto do Gin e busca um usuário por ID
func (uc userControllerInterface) FindUserByID(c *gin.Context) { 
	logger.Info("Initiating findUserByID controller",
		zap.String("journey", "findUserByID"), 
	)

	userId := c.Param("userId") // Obtém o ID do usuário da requisição

	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		logger.Error("Error trying to validate userId", err, 
			zap.String("journey", "findUserByID"),
		)

		errorMessage := rest_err.NewBadRequestError( 
			"UserID is not a valid id", 
		)

		c.JSON(errorMessage.Code, errorMessage) 
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId) // Busca o usuário por ID
	if err != nil {                                          
		logger.Error("Error trying to call findUserByID services", err,
			zap.String("journey", "findUserById"), 
		)

		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("FindUserByID controller executed successfully", 
		zap.String("journey", "findUserById"), 
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain, // Domínio de usuário
	)) // Retorna o usuário encontrado com sucesso

}

 // Função que recebe um contexto do Gin e busca um usuário por email
func (uc userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Initiating findUserByEmail controller", 
		zap.String("journey", "findUserByEmail"), 
	)

	userEmail := c.Param("userEmail") // Obtém o email do usuário da requisição

	if _, err := mail.ParseAddress(userEmail); err != nil { // Verifica se o email do usuário é válido
		logger.Error("Error trying to validate userEmail", err, 
			zap.String("journey", "findUserByEmail"), 
		)

		errorMessage := rest_err.NewBadRequestError( 
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage) 
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail) // Busca o usuário por email
	if err != nil {                                                 
		logger.Error("Error trying to call findUserByEmail services", err, 
			zap.String("journey", "findUserByEmail"), 
		)

		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("findUserByEmail controller executed successfully", 
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
