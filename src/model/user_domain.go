package model

// Struct que define o domínio de usuário
type userDomain struct {
	ID       string // ID do usuário
	email    string // Email do usuário
	password string // Senha do usuário
	name     string // Nome do usuário
	age      int8   // Idade do usuário
}

// Função que retorna o ID do usuário
func (ud *userDomain) GetID() string { 
	return ud.ID 
}

// Função que define o ID do usuário
func (ud *userDomain) SetID(id string) { 
	ud.ID = id 
}

// Função que retorna o email do usuário
func (ud *userDomain) GetEmail() string { 
	return ud.email 
}

// Função que retorna a senha do usuário
func (ud *userDomain) GetPassword() string { 
	return ud.password 
}

// Função que retorna o nome do usuário
func (ud *userDomain) GetName() string { 
	return ud.name 
}

// Função que retorna a idade do usuário
func (ud *userDomain) GetAge() int8 { 
	return ud.age 
}
