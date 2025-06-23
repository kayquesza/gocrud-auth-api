package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/logger"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Função que recebe um contexto do Gin e deleta um usuário
func (uc userControllerInterface) DeleteUser(c *gin.Context) { 
	logger.Info("Init deleteUser Controller", 
		zap.String("journey", "deleteUser"), 
	)

	userId := c.Param("userId")                                  // Obtém o ID do usuário da requisição
	if _, err := primitive.ObjectIDFromHex(userId); err != nil { // Verifica se o ID do usuário é válido
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value") 
		c.JSON(errRest.Code, errRest)                                                 
		return
	}

	err := uc.service.DeleteUser(userId) // Deleta o usuário
	if err != nil {                      
		logger.Error("Error trying to call deleteUser service.", err, 
			zap.String("journey", "deleteUser")) 
		c.JSON(err.Code, err) // Retorna o erro 500 com a mensagem de erro
		return
	}

	logger.Info("User deleted successfully", 
		zap.String("userId", userId),        
		zap.String("journey", "deleteUser")) 

	c.Status(http.StatusOK) // Retorna o status 200 OK
}
