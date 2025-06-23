package request

type UserLogin struct { // Estrutura para representar o login de um usuário
	Email    string `json:"email" binding:"required,email"`                        // Email do usuário
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"` // Senha do usuário
}
