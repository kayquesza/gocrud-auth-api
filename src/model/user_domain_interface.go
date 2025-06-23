package model

import "github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"

// Interface que define os métodos que o domínio de usuário deve implementar
type UserDomainInterface interface {
	GetEmail() string    // Retorna o email do usuário
	GetPassword() string // Retorna a senha do usuário
	GetName() string     // Retorna o nome do usuário
	GetAge() int8        // Retorna a idade do usuário
	GetID() string       // Retorna o ID do usuário

	SetID(string) // Define o ID do usuário

	EncryptPassword() error                     // Criptografa a senha do usuário
	GenerateToken() (string, *rest_err.RestErr) // Gera o token do usuário
}

// Função que cria um novo domínio de usuário
func NewUserDomain(
	email, password, name string, // Email, senha e nome do usuário
	age int8, // Idade do usuário
) UserDomainInterface { // Retorno de uma interface com os métodos
	return &userDomain{ // Retorna um domínio de usuário
		email:    email,    // Email do usuário
		password: password, // Senha do usuário
		name:     name,     // Nome do usuário
		age:      age,      // Idade do usuário
	}
}

// Função que cria um novo domínio de login de usuário
func NewUserLoginDomain(
	email, password string, // Email e senha do usuário
) UserDomainInterface { // Retorno de uma interface com os métodos
	return &userDomain{ // Retorna um domínio de login de usuário
		email:    email,    // Email do usuário
		password: password, // Senha do usuário
	}
}

// Função que cria um novo domínio de atualização de usuário
func NewUserUpdateDomain(
	name string, // Nome do usuário
	age int8, // Idade do usuário
) UserDomainInterface { // Retorno de uma interface com os métodos
	return &userDomain{ // Retorna um domínio de atualização de usuário
		name: name, // Nome do usuário
		age:  age,  // Idade do usuário
	}
}
