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

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreatUser Controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	// Pegará o body da requisição e irá fazer o bind para a struct UserRequest
	// Se houver algum erro, retornará um erro 400 com a mensagem de erro
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidadeUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserService(domain)
	if err != nil {
		logger.Error("Error trying to call CreatUser service.", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succesfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
