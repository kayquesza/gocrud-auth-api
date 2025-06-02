package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"` // Informar o erro ao usuário
	Err     string   `json:"error"`   // Mostrar o significado do erro na requisição
	Code    int      `json:"code"`    // Qual o código da requisição retornada ao usuário
	Causes  []Causes `json:"causes"`  // Quais as causas de erro dentro da aplicação
}

// Uma lista de campos incorretos dentro da aplicação
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Satisfazer a interface de erro do Go
func (r *RestErr) Error() string {
	return r.Message
}

// Construtor para o objeto
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request", // Status HTTP 400 - Requisição malformada (Bad Request)
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request", // Status HTTP 400 - Requisição malformada (Bad Request)
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server Error", // Status HTTP 500 - Erro interno do servidor (Internal Server Error)
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not_Found", // Status HTTP 404 - Recurso não encontrado (Not Found)
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden", // Status HTTP 403 - Acesso proibido (Forbidden)
		Code:    http.StatusForbidden,
	}
}
