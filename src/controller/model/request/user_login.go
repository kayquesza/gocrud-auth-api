package request

// Estrutura para representar o login de um usuário
type UserLogin struct { 
	Email    string `json:"email" binding:"required,email"`                        
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"` 
}
