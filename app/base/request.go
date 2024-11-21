package base

import (
	"fmt"
	"golang/constants"

	"github.com/go-playground/validator/v10"

	customValidator "golang/pkg/utils/validator"
)

var validate *validator.Validate

type Validator interface {
	Validate() map[string]string
}

func init() {
	validate = validator.New()

	validate.RegisterValidation("email_custom", customValidator.IsValidEmail)
}

func ValidateStruct(req Validator) map[string]string {
	err := validate.Struct(req)
	if err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			if messageTemplate, ok := constants.Config.AppConfig.ErrorsMessages[err.Tag()]; ok {
				errors[err.Field()] = fmt.Sprintf(messageTemplate, err.Field())
			} else {
				errors[err.Field()] = fmt.Sprintf("Validation failed for field '%s'", err.Field())
			}
		}

		return errors
	}

	return nil
}
