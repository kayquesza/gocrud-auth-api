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

func (uc userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Initiating findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err,
			zap.String("journey", "findUserByID"),
		)

		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err,
			zap.String("journey", "findUserById"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}

func (uc userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Initiating findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err,
			zap.String("journey", "findUserByEmail"),
		)

		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services", err,
			zap.String("journey", "findUserByEmail"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("findUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
