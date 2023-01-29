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

// Validation Nested Struct
func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Id     string  `validate:"required"`
		Name   string  `validate:"required"`
		Addres Address `validate:"required"`
	}

	request := User{
		Id:   "",
		Name: "",
		Addres: Address{
			City:    "",
			Country: "",
		},
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation Collection
func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	request := User{
		Id:   "01",
		Name: "jalal",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
		},
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestBacisCollection(t *testing.T) {
	type User struct {
		Id      string   `validate:"required"`
		Name    string   `validate:"required"`
		Hobbies []string `validate:"dive,required,min=3"`
	}
	user := User{
		Id:      "01",
		Name:    "jalal",
		Hobbies: []string{},
	}
	validation := validator.New()
	err := validation.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation Map
func TestMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}
	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Hobbies []string          `validate:"dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=3,endkeys,dive"`
	}
	user := User{
		Id:      "01",
		Name:    "Jalal",
		Hobbies: []string{"Gaming", "Coding"},
		Schools: map[string]School{
			"SD": {
				Name: "",
			},
			"SMP": {
				Name: "MTsN 1 ",
			},
		},
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}
	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Hobbies []string          `validate:"dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=3,endkeys,dive"`
		Wallet  map[string]int    `validate:"dive,keys,required,endkeys,required,gt=20000"`
	}
	user := User{
		Id:      "01",
		Name:    "Jalal",
		Hobbies: []string{"Gaming", "Coding"},
		Schools: map[string]School{
			"SDN": {
				Name: "SDN 01",
			},
			"SMP": {
				Name: "MTsN 1 ",
			},
		},
		Wallet: map[string]int{
			"BNI": 10000,
		},
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Alias Tag
func TestAlias(t *testing.T) {
	validation := validator.New()
	validation.RegisterAlias("var", "required,max=10")

	type Seller struct {
		Id     string `validate:"var"`
		Name   string `validate:"var"`
		Owner  string `validate:"var"`
		Slogan string `validate:"var"`
	}
	seller := Seller{
		Id:     "",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}
	err := validation.Struct(seller)

	if err != nil {
		fmt.Println(err.Error())
	}
}
