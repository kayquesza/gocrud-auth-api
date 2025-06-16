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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser Controller",
		zap.String("journey", "loginUser"),
	)

	var userRequest request.UserLogin

	

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "loginUser"))
		errRest := validation.ValidadeUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("Error trying to call loginUser service.", err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succesfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))

}
