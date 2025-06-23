package view

import (
	"github.com/kayquesza/gocrud-auth-api/src/controller/model/response"
	"github.com/kayquesza/gocrud-auth-api/src/model"
)

// Função que converte um domínio de usuário para uma resposta
func ConvertDomainToResponse(
	userDomain model.UserDomainInterface, // Domínio de usuário
) response.UserResponse { // Retorna uma resposta de usuário
	return response.UserResponse{ // Retorna uma resposta de usuário
		ID:    userDomain.GetID(),    // ID do usuário
		Email: userDomain.GetEmail(), // Email do usuário
		Name:  userDomain.GetName(),  // Nome do usuário
		Age:   userDomain.GetAge(),   // Idade do usuário
	}
}
