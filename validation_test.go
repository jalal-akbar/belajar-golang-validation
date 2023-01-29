package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

// Validation
func TestValidation(t *testing.T) {
	validate := validator.New()

	if validate == nil {
		t.Error("Validate is nil")
	}

}

// Validation Variable
func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "jalal"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}

}

// Validation Two Variable
func TestValidationTwoVariable(t *testing.T) {
	validate := validator.New()
	password := "rahasia"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}

}

// Validation Multiple Tag
func TestValidationMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "jalal"

	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation With Tag Parameter
func TestValidationTagParameter(t *testing.T) {
	validate := validator.New()
	res := "99"

	err := validate.Var(res, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

// Validation Struct
func TestValidationStruct(t *testing.T) {

	validate := validator.New()
	user := LoginRequest{"muh@gmail.com", "muhakbar"}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation Error
func TestValidationError(t *testing.T) {
	validate := validator.New()
	user := LoginRequest{
		Username: "muh",
		Password: "muh",
	}
	err := validate.Struct(user)

	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "With error", fieldError.Error())
		}
	}
}

// Validation Cross Field

type CrossField struct {
	Username        string `validate:"required,email"`
	Password        string `validate:"required,min=5"`
	ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
}

func TestValidationCrossField(t *testing.T) {
	crossField := CrossField{
		Username:        "jalal@gmail.com",
		Password:        "benar",
		ConfirmPassword: "benar",
	}
	validate := validator.New()
	err := validate.Struct(crossField)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
