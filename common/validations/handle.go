package validations

import (
	"fmt"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/go-playground/validator/v10"
)

func debugFieldError(v validator.FieldError) {
	fmt.Println("ini field", v.Field())
	fmt.Println("ini kind", v.Kind())
	fmt.Println("ini actual tag", v.ActualTag())
	fmt.Println("ini tag", v.Tag())
	fmt.Println("ini value", v.Value())
	fmt.Println("ini type", v.Type())
	fmt.Println("ini namespace", v.Namespace())
	fmt.Println("ini param", v.Param())
	fmt.Println("ini struct field", v.StructField())
	fmt.Println("ini struct namespace", v.StructNamespace())
}

func HandleValidationErrors(errs validator.ValidationErrors) (res model.ValidationErrorWebServiceResponse) {
	res.Status = http.StatusUnprocessableEntity
	finalErrors := make(map[string]string)

	errorMessage := "Terjadi kesalahan pada validasi pada field:"

	for _, v := range errs {
		translateFunction, ok := customErrors[v.Tag()]
		if ok {
			translatedFieldName, found := customMessages[v.Field()]
			if !found {
				translatedFieldName = v.Field()
			}

			// Lowercase the field name for the JSON key convention
			fieldKey := v.Field()
			if len(fieldKey) > 0 {
				fieldKey = string(fieldKey[0]|32) + fieldKey[1:] // simple uncapitalize
			}

			msg := translateFunction(v, translatedFieldName)
			finalErrors[fieldKey] = msg
		}

		errorMessage += " " + v.Field()
	}

	res.Message = errorMessage
	res.DetailErrors = finalErrors
	res.Data = nil
	return res
}
