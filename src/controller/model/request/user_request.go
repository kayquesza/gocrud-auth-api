package request

// Estrutura para representar o cadastro de um usuário
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`                        // Email do usuário
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"` // Senha do usuário
	Name     string `json:"name" binding:"required,min=3,max=50"`                  // Nome do usuário
	Age      int8   `json:"age" binding:"required,min=18,max=120"`                 // Idade do usuário
}

// Estrutura para representar a atualização de um usuário
type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=50"`  // Nome do usuário
	Age  int8   `json:"age" binding:"omitempty,min=18,max=120"` // Idade do usuário
}
