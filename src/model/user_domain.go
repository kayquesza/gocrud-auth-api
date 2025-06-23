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
func (ud *userDomain) GetID() string { // Retorna o ID do usuário
	return ud.ID // Retorna o ID do usuário
}

// Função que define o ID do usuário
func (ud *userDomain) SetID(id string) { // Define o ID do usuário
	ud.ID = id // Define o ID do usuário
}

// Função que retorna o email do usuário
func (ud *userDomain) GetEmail() string { // Retorna o email do usuário
	return ud.email // Retorna o email do usuário
}

// Função que retorna a senha do usuário
func (ud *userDomain) GetPassword() string { // Retorna a senha do usuário
	return ud.password // Retorna a senha do usuário
}

// Função que retorna o nome do usuário
func (ud *userDomain) GetName() string { // Retorna o nome do usuário
	return ud.name // Retorna o nome do usuário
}

// Função que retorna a idade do usuário
func (ud *userDomain) GetAge() int8 { // Retorna a idade do usuário
	return ud.age // Retorna a idade do usuário
}
