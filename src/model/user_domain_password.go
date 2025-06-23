package model

import "golang.org/x/crypto/bcrypt"

// Função que criptografa a senha do usuário
func (ud *userDomain) EncryptPassword() error { // Retorna um erro
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost) // Gera a senha criptografada, possível definiir o custo da criptografia
	if err != nil {                                                                             // Se houver algum erro, retorna um erro
		return err // Retorna um erro
	}
	ud.password = string(hashedPassword) // Converte a senha criptografada para uma string
	return nil                           // Retorna nil
} // Implementa o método EncryptPassword que irá criptografar a senha do usuário utilizando o pacote bcrypt antes de salvar no banco de dados
