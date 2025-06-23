package view

import (
	"github.com/kayquesza/gocrud-auth-api/src/controller/model/response"
	"github.com/kayquesza/gocrud-auth-api/src/model"
)

// Função que recebe o domínio do usuário e retorna uma struct para uma resposta
func ConvertDomainToResponse(
	userDomain model.UserDomainInterface, // Recebe um objeto do domínio do usuário
) response.UserResponse { 
	return response.UserResponse{ 
		ID:    userDomain.GetID(),    
		Email: userDomain.GetEmail(), 
		Name:  userDomain.GetName(),  
		Age:   userDomain.GetAge(),   
	}
}
