package validator

import (
	"fmt"
	"golang/constants"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

// IsValidEmail checks if the provided email is valid
func IsValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Example: reject specific email
	if email == "dinhtq.dev@gmail.com" {
		return false
	}

	// Regular expression to validate email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return re.MatchString(email)
}

// IsEmpty checks if the provided string is empty or contains only whitespace
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// IsNotEmpty checks if the provided string is not empty
func extractFieldFromDuplicateKeyMessage(message string) string {
	re := regexp.MustCompile(`for key '([^']*)'`)
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		// Map the index name to the actual field name if needed
		indexName := matches[1]
		return mapIndexNameToFieldName(indexName)
	}
	return ""
}

func extractFieldFromCannotBeNullMessage(message string) string {
	re := regexp.MustCompile(`Column '([^']*)' cannot be null`)
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func mapIndexNameToFieldName(indexName string) string {
	indexToFieldMap := map[string]string{
		"uni_users_email": "email",
		"uni_users_name":  "name",
	}
	if field, exists := indexToFieldMap[indexName]; exists {
		return field
	}
	return indexName
}

// HandleMySQLError processes MySQL errors and returns a map of custom error messages
// based on the specific MySQL error number.
// 
// Params:
//   - err: The MySQL error to be processed (of type error).
//
// Returns:
//   - A map of string keys and values where the key is the field or type of error,
//     and the value is the corresponding error message.
//   - An HTTP status code representing the type of error (e.g., 400 for Bad Request).
func HandleMySQLError(err error) (map[string]string, int) {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		errors := make(map[string]string)

		switch mysqlErr.Number {
		case 1062: // Duplicate entry
			// Parse the error message to extract the field causing the error
			field := extractFieldFromDuplicateKeyMessage(mysqlErr.Message)
			if field != "" {
				errors[field] = fmt.Sprintf("The %s field must be unique.", field)
			} else {
				errors["general"] = "A duplicate entry exists in the database."
			}
			return errors, http.StatusBadRequest

		case 1048: // Column cannot be null
			field := extractFieldFromCannotBeNullMessage(mysqlErr.Message)
			if field != "" {
				errors[field] = fmt.Sprintf("The %s field is required.", field)
			} else {
				errors["general"] = "A required field is missing."
			}
			return errors, http.StatusBadRequest

		case 1451: // Foreign key constraint fails (Cannot delete parent)
			errors["general"] = "Cannot delete or update because of foreign key constraints."
			return errors, http.StatusConflict

		case 1452: // Foreign key constraint fails (Cannot add or update child)
			errors["general"] = "Cannot add or update due to foreign key constraints."
			return errors, http.StatusConflict

		case 1216, 1217: // Foreign key constraint failure
			errors["general"] = "Foreign key constraint failure. Please check related records."
			return errors, http.StatusConflict

		case 1366: // Invalid string value
			errors["general"] = "Invalid value provided for one of the fields."
			return errors, http.StatusBadRequest

		default:
			errors["general"] = "A database error occurred."
			return errors, http.StatusInternalServerError
		}
	}

	return map[string]string{"general": "An unknown error occurred."}, http.StatusInternalServerError
}

// ValidateStruct validates the struct passed as an argument using the go-playground/validator package.
// It checks the fields of the struct against the specified validation tags and returns a map of error messages
// for any failed validation checks.
//
// Params:
//   - v: The struct to be validated. This can be any struct that has validation tags defined on its fields.
//
// Returns:
//   - A map of string keys and values where the key is the name of the field that failed validation,
//     and the value is the corresponding error message. If there are no validation errors, it returns nil.
func ValidateStruct(v interface{}) map[string]string {
    validate := validator.New()
    err := validate.Struct(v)
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
