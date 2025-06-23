package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc userControllerInterface) DeleteUser(c *gin.Context) { // Função que recebe um contexto do Gin e deleta um usuário
	logger.Info("Init deleteUser Controller", // Registra o início da deleção de um usuário
		zap.String("journey", "deleteUser"), // Registra a jornada da deleção de um usuário
	)

	userId := c.Param("userId")                                  // Obtém o ID do usuário da requisição
	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value") // Cria um erro de requisição malformada
		c.JSON(errRest.Code, errRest)                                                 // Retorna o erro de requisição malformada
		return
	}

	err := uc.service.DeleteUser(userId) // Deleta o usuário
	if err != nil {                      // Se houver algum erro, retornará um erro 500 com a mensagem de erro
		logger.Error("Error trying to call deleteUser service.", err, // Registra o erro de deleção do usuário
			zap.String("journey", "deleteUser")) // Registra a jornada da deleção de um usuário
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User deleted successfully", // Registra o sucesso da deleção do usuário
		zap.String("userId", userId),        // Registra o ID do usuário
		zap.String("journey", "deleteUser")) // Registra a jornada da deleção de um usuário

	c.Status(http.StatusOK) // Retorna o status 200 OK
}
