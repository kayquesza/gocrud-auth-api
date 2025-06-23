package converter

import (
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity"
)

// Função que converte um domínio de usuário para uma entidade de usuário
func ConvertDomainToEntity(
	domain model.UserDomainInterface, // Domínio de usuário
) *entity.UserEntity {
	return &entity.UserEntity{ // Retorna uma entidade de usuário
		Email:    domain.GetEmail(),    // Email do usuário
		Password: domain.GetPassword(), // Senha do usuário
		Name:     domain.GetName(),     // Nome do usuário
		Age:      domain.GetAge(),      // Idade do usuário
	}
}
