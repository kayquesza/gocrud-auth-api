package view

import (
	"github.com/kayquesza/gocrud-auth-api/src/controller/model/response"
	"github.com/kayquesza/gocrud-auth-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
