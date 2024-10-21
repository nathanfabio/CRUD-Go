package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine(). (*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
		}
}


func ValidateUserError (
	validation_err error,
	) *errs.Errs {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return errs.NewBadRequestErrs("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []errs.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := errs.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return errs.NewBadRequestValidationErrs("Invalid fields", errorsCauses)
	} else {
		return errs.NewBadRequestErrs("Error to convert fields")
	}
}