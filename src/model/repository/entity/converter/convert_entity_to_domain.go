package converter

import (
	"github.com/kayquesza/gocrud-auth-api/src/model"
	"github.com/kayquesza/gocrud-auth-api/src/model/repository/entity"
)

// Função que converte uma entidade de usuário para um domínio de usuário
func ConvertEntityToDomain(
	entity entity.UserEntity, // Entidade de usuário
) model.UserDomainInterface {
	domain := model.NewUserDomain( // Cria um novo domínio de usuário
		entity.Email,    // Email do usuário
		entity.Password, // Senha do usuário
		entity.Name,     // Nome do usuário
		entity.Age,      // Idade do usuário
	)

	domain.SetID(entity.ID.Hex()) // Hex retorna somente o valor
	return domain                 // Retorna o domínio de usuário
}
