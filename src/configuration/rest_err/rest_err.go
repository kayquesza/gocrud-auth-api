package rest_err

import "net/http"

type RestErr struct { // Estrutura para representar erros REST
	Message string   `json:"message"` // Informar o erro ao usuário
	Err     string   `json:"error"`   // Mostrar o tipo do erro na requisição
	Code    int      `json:"code"`    // Qual o código da requisição retornada ao usuário
	Causes  []Causes `json:"causes"`  // Quais as causas de erro dentro da aplicação
}

// Uma lista de campos incorretos dentro da aplicação
type Causes struct { // Estrutura para detalhar causas de erro
	Field   string `json:"field"`   // Qual o campo que está incorreto
	Message string `json:"message"` // Qual a mensagem do específica do erro
}

// Satisfazer a interface de erro do Go
func (r *RestErr) Error() string { // Função que retorna a mensagem de erro
	return r.Message // Retorna a mensagem de erro
}

// Construtor que cria um erro rest
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr { // Função que cria um erro rest para uma requisição malformada
	return &RestErr{
		Message: message,
		Err:     "Bad Request", // Status HTTP 400 - Requisição malformada (Bad Request)
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr { // Função que cria um erro rest para uma requisição não autorizada
	return &RestErr{
		Message: message,
		Err:     "Unauthorized", // Status HTTP 401
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr { // Função que cria um erro rest para uma requisição malformada
	return &RestErr{
		Message: message,
		Err:     "Bad Request", // Status HTTP 400 - Requisição malformada (Bad Request)
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr { // Função que cria um erro rest para um erro interno do servidor
	return &RestErr{
		Message: message,
		Err:     "Internal Server Error", // Status HTTP 500 - Erro interno do servidor (Internal Server Error)
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr { // Função que cria um erro rest para um recurso não encontrado
	return &RestErr{
		Message: message,
		Err:     "Not_Found", // Status HTTP 404 - Recurso não encontrado (Not Found)
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr { // Função que cria um erro rest para um acesso proibido
	return &RestErr{
		Message: message,
		Err:     "Forbidden", // Status HTTP 403 - Acesso proibido (Forbidden)
		Code:    http.StatusForbidden,
	}
}
