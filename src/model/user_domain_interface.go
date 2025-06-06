package model

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
