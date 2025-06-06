package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string

	SetID(string)

	EncryptPassword() error
}

func NewUserDomain(
	email, password, name string,
	age int8,
) *userDomain { // Retorno de uma interface com os métodos
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) GetID() string {
	return ud.ID // Retorna o ID do usuário
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.password = string(hashedPassword)
	return nil
} // Implementa o método EncryptPassword que irá criptografar a senha do usuário utilizando o pacote bcrypt antes de salvar no banco de dados
