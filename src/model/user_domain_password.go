package model

import "golang.org/x/crypto/bcrypt"

// Função que criptografa a senha do usuário
func (ud *userDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost) // Gera a senha criptografada, possível definir o custo da criptografia
	if err != nil {
		return err // Retorna um erro
	}
	ud.password = string(hashedPassword) // Converte a senha criptografada para uma string
	return nil
} // Implementa o método EncryptPassword que irá criptografar a senha do usuário utilizando o pacote bcrypt antes de salvar no banco de dados
