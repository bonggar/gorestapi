package helper

import (
	"unicode"

	"gopkg.in/go-playground/validator.v9"
)

//ConstructErrors : Function to construct errors to be used in response
func ConstructErrors(err error) []ErrorField {
	errorList, success := err.(validator.ValidationErrors)
	errorFields := []ErrorField{}
	if success {
		for _, value := range errorList {
			fieldIDString := SetFirstLetterToLowerCase(value.StructField())
			fieldValue := value.Value().(string)
			fieldErrorCaused := value.Translate(nil)
			fieldErrorMessage := "Invalid Value"
			switch value.Tag() {
			case "required":
				fieldErrorMessage = "Required"
			case "alpha":
				fieldErrorMessage = "Only letters allowed"
			case "email":
				fieldErrorMessage = "Invalid email format"
			case "numeric":
				fieldErrorMessage = "Only numbers allowed"
			case "min":
				fieldErrorMessage = "Minumum Character " + value.Param()
			case "max":
				fieldErrorMessage = "Maximum Character " + value.Param()
			}
			errorFields = append(errorFields, ErrorField{
				fieldIDString, fieldValue, fieldErrorCaused, fieldErrorMessage})
		}
	}
	return errorFields
}

//SetFirstLetterToLowerCase : Function to set first letter of string to lower case.
//This function exist because struct field name must start with capital letter
//This function is useful so the front end can detect which field has an error
func SetFirstLetterToLowerCase(toBeConvertedString string) string {
	fieldNameBytes := []rune(toBeConvertedString)
	fieldNameBytes[0] = unicode.ToLower(fieldNameBytes[0])
	return string(fieldNameBytes)
}
