package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// Custom Validation
func toUpperValidation(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}
func TestCustomValidation(t *testing.T) {
	validation := validator.New()
	validation.RegisterValidation("toupper", toUpperValidation)

	type Seller struct {
		Id     string `validate:"required,toupper"`
		Name   string `validate:"required"`
		Owner  string `validate:"required"`
		Slogan string `validate:"required"`
	}
	seller := Seller{
		Id:     "akbar",
		Name:   "akbar",
		Owner:  "akbar",
		Slogan: "akbar",
	}
	err := validation.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Custom Validation Parameter
var regexNumber = regexp.MustCompile("[0-9]")

func mustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}
	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}
	return len(value) == length
}
func TestCustomValidationParameter(t *testing.T) {
	validation := validator.New()
	validation.RegisterValidation("pin", mustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}
	login := Login{
		Phone: "0810238032",
		Pin:   "23456",
	}

	err := validation.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Or Rule
func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	login := Login{
		Username: "muh",
		Password: "24234",
	}
	validation := validator.New()
	err := validation.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Custom Validation Cross Field
func mustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}
	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())
	return firstValue == secondValue
}

func TestCustomValidationCrossField(t *testing.T) {
	validation := validator.New()
	validation.RegisterValidation("fields_equals_ignore_case", mustEqualIgnoreCase)

	type Login struct {
		Username string `validate:"required,fields_equals_ignore_case=Email|fields_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
		Phone    string `validate:"required,numeric"`
	}
	login := Login{
		Username: "jalal",
		Email:    "muh@gmail.com",
		Password: "jalal08493",
		Phone:    "08138163449",
	}

	err := validation.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Struct Level Validation

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func musyValidRegisterStruct(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// sukses
	} else {
		// gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validation := validator.New()
	validation.RegisterStructValidation(musyValidRegisterStruct, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "jalal",
		Email:    "jalal@gmail.com",
		Phone:    "08127772424",
		Password: "jalal4597",
	}

	err := validation.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
