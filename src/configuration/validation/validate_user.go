package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding" // usado para vincular dados JSON (do corpo da requisição) a uma struct
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
)

// Variáveis de ambiente para a validação de usuário
var (
	Validate = validator.New() // Cria uma nova instância do validador
	transl   ut.Translator     // Responsável por traduzir mensagens de validação
)

func init() { // Função para inicializar a validação de usuário

	// Método para obter o validador do binding
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok { // Mecanismo de validação padrão do Gin
		en := en.New() // Configura o validador para o idioma inglês
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en") // Cria um novo tradutor com o idioma inglês
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidadeUserError( // Função que recebe um erro de validação e retorna um erro padronizado
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError               // Variável para capturar o erro de desempacotamento JSON
	var jsonValidationError validator.ValidationErrors // Variável para armazenar o erro de validação

	if errors.As(validation_err, &jsonErr) { // Verifica se o erro é um erro de desempacotamento JSON
		return rest_err.NewBadRequestError("Invalid field type") // Retorna um erro de requisição malformada
	} else if errors.As(validation_err, &jsonValidationError) { // Verifica se o erro é um erro de validação
		errorCauses := []rest_err.Causes{} // Variável para armazenar as causas do erro

		for _, e := range validation_err.(validator.ValidationErrors) { // Itera sobre os erros de validação
			cause := rest_err.Causes{ // Cria uma causa do erro
				Message: e.Translate(transl), // Traduz a mensagem do erro
				Field:   e.Field(),           // Obtém o campo que causou o erro
			}
			errorCauses = append(errorCauses, cause) // Adiciona a causa do erro à lista de causas
		}
		return rest_err.NewBadRequestValidationError("Some fileds are invalid", errorCauses) // Retorna um erro de requisição malformada
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields") // Se o erro não for um erro de desempacotamento JSON ou um erro de validação, retorna um erro genérico
	}

}
