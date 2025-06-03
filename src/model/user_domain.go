package model

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}

// Sem o uso de tags, pois o Domain não pode ser "exportável"
type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.Password = string(hashedPassword)
	return nil
} // Implementa o método EncryptPassword que irá criptografar a senha do usuário utilizando o pacote bcrypt antes de salvar no banco de dados

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr                    // Recebe um objeto do usuário e retorna um erro se houver
	UpdateUser(string) *rest_err.RestErr              // Recebe uma string com o ID do usuário e um objeto do usuário para atualizar
	FindUser(string) (*UserDomain, *rest_err.RestErr) // Recebe uma string com o ID do usuário e retorna um objeto do usuário e um erro se houver
	DeleteUser(string) *rest_err.RestErr              // Recebe uma string com o ID do usuário e retorna um erro se houver
}
