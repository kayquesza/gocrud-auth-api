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

func (uc userControllerInterface) FindUserByID(c *gin.Context) { // Função que recebe um contexto do Gin e busca um usuário por ID
	logger.Info("Initiating findUserByID controller", // Mensagem de log
		zap.String("journey", "findUserByID"), // Jornada da busca de um usuário por ID
	)

	userId := c.Param("userId") // Obtém o ID do usuário da requisição

	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		logger.Error("Error trying to validate userId", err, // Mensagem de log
			zap.String("journey", "findUserByID"), // Jornada da busca de um usuário por ID
		)

		errorMessage := rest_err.NewBadRequestError( // Cria um erro de requisição malformada
			"UserID is not a valid id", // Mensagem de erro
		)

		c.JSON(errorMessage.Code, errorMessage) // Retorna o erro de requisição malformada
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId) // Busca o usuário por ID
	if err != nil {                                            // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call findUserByID services", err, // Mensagem de log
			zap.String("journey", "findUserById"), // Jornada da busca de um usuário por ID
		)

		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("FindUserByID controller executed successfully", // Mensagem de log
		zap.String("journey", "findUserById"), // Jornada da busca de um usuário por ID
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse( // Converte o domínio de usuário para uma resposta
		userDomain, // Domínio de usuário
	)) // Retorna o usuário encontrado com sucesso

}

func (uc userControllerInterface) FindUserByEmail(c *gin.Context) { // Função que recebe um contexto do Gin e busca um usuário por email
	logger.Info("Initiating findUserByEmail controller", // Mensagem de log
		zap.String("journey", "findUserByEmail"), // Jornada da busca de um usuário por email
	)

	userEmail := c.Param("userEmail") // Obtém o email do usuário da requisição

	if _, err := mail.ParseAddress(userEmail); err != nil { // Verifica se o email do usuário é válido
		logger.Error("Error trying to validate userEmail", err, // Mensagem de log
			zap.String("journey", "findUserByEmail"), // Jornada da busca de um usuário por email
		)

		errorMessage := rest_err.NewBadRequestError( // Cria um erro de requisição malformada
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage) // Retorna o erro de requisição malformada
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail) // Busca o usuário por email
	if err != nil {                                                  // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call findUserByEmail services", err, // Mensagem de log
			zap.String("journey", "findUserByEmail"), // Jornada da busca de um usuário por email
		)

		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("findUserByEmail controller executed successfully", // Mensagem de log
		zap.String("journey", "findUserByEmail"), // Jornada da busca de um usuário por email
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
