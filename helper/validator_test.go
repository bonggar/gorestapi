package helper

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name    string `validate:"min=3,alpha"`
	Address string `validate:"required"`
}

func TestConstructErrorsUnmatchedErrorFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	err := errors.New("math: square root of negative number")
	errorFields := ConstructErrors(err)
	if len(errorFields) == 0 {
		t.Logf("Prevent Panic when parsing failed")
	} else {
		t.Fail()
	}
}

func TestConstructErrorsWithValidatorTag(t *testing.T) {
	gin.SetMode(gin.TestMode)
	validate := validator.New()

	str1 := new(string)
	*str1 = "1"

	str2 := new(string)
	*str2 = ""
	err := validate.Struct(User{Name: *str1})
	errorFields := ConstructErrors(err)
	if len(errorFields) > 0 {
		t.Logf("Convert validator error to JSON format as response \n")
		bytesJSON, _ := json.MarshalIndent(errorFields, "", " ")
		t.Logf(string(bytesJSON))
	} else {
		t.Fail()
	}
}

func TestConvertFirstLetterToLowerCase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	str := "Test"
	convertedStr := SetFirstLetterToLowerCase(str)
	if convertedStr == "test" {
		t.Logf("Converting First Letter to Lowercase Success")
	} else {
		t.Fail()
	}
}
