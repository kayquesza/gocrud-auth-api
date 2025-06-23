package response

// Estrutura para representar a resposta de dados de um usuário, sem conter a senha
type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
