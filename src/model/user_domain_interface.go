package model

import "github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string

	SetID(string)

	EncryptPassword() error
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface { // Retorno de uma interface com os métodos
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface { // Retorno de uma interface com os métodos
	return &userDomain{
		name: name,
		age:  age,
	}
}
