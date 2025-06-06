package model

type userDomain struct {
	ID       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetID() string {
	return ud.ID // Retorna o ID do usuário
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
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
