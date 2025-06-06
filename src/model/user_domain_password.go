package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.password = string(hashedPassword)
	return nil
} // Implementa o método EncryptPassword que irá criptografar a senha do usuário utilizando o pacote bcrypt antes de salvar no banco de dados
