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

var (
	Validate = validator.New() // Cria uma nova instância do validador
	transl   ut.Translator     // Responsável por traduzir mensagens de validação
)

func init() {

	// Método para obter o validador do binding
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok { // Mecanismo de validação padrão do Gin
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en") // Cria um novo tradutor com o idioma inglês
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidadeUserError( // Transforma os erros de validação em um objeto de erro customizado
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fileds are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
