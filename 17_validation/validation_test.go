package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	request := ""

	err := validate.Var(request, "required")
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()
	newPassword := "Svelte"
	confirmNewPassword := "Svelte"

	err := validate.VarWithValue(newPassword, confirmNewPassword, "eqfield")
	if err != nil {
		fmt.Println(err)
	}
}

func TestMultipleTagValidation(t *testing.T) {
	validate := validator.New()
	request := "123"

	// err := validate.Var(request, "required,number")
	err := validate.Var(request, "required,numeric")
	if err != nil {
		fmt.Println(err)
	}
}

func TestTagarameter(t *testing.T) {
	validate := validator.New()
	request := "10000"

	err := validate.Var(request, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err)
	}
}

func TestStructValidation(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email,max=255"`
		Password string `validate:"required,min=8,max=255"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "x@x.x",
		Password: "Svelte78",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email,max=255"`
		Password string `validate:"required,min=8,max=255"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "x@x.",
		Password: "Svelte",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error:", fieldError.Field(), "- on tag:", fieldError.Tag(), "- error details:", fieldError)
			fmt.Println("actual tag:", fieldError.ActualTag())
		}
	}
}

func TestValidationCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,max=255"`
		Password        string `validate:"required,min=8,max=255"`
		ConfirmPassword string `validate:"required,min=8,max=255,eqfield=Password"`
	}

	validate := validator.New()
	registerUser := RegisterUser{
		Username:        "Svelte",
		Password:        "Svelte Kit",
		ConfirmPassword: "Svelte Kit",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err)
	}
}

func TestNestedStructValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"` // tag di fieldnya jika dihapus tidak masalah
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestBasicCollectionValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{"Gaming", "Ss", "Coding", "Xx", ""},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestMapValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required,min=10"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{"Gaming", "Ss", "Coding", "Xx", ""},
		Schools: map[string]School{
			"first":  {Name: "first school"},
			"second": {Name: "second"},
			"":       {Name: ""},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestBasicMapValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required,min=10"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys,required"`
		Wallets   map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=1000"`
	}

	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{City: "", Country: ""},
			{City: "", Country: ""},
			{City: "", Country: ""},
		},
		Hobbies: []string{"Gaming", "Ss", "Coding", "Xx", ""},
		Schools: map[string]School{
			"first":  {Name: "first school"},
			"second": {Name: "second"},
			"":       {Name: ""},
		},
		Wallets: map[string]int{
			"BANK1": 1000000,
			"BANK2": 0,
			"":      0,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:   "S0001",
		Name: "Svelte Kit",
		// Owner:  "Svelte",
		// Slogan: "Fastest frontend library/framework",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err)
	}
}

func mustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
		return true
	} else {
		return false
	}
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", mustValidUsername)

	type User struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := User{
		Username: "SVEL",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

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
	validate := validator.New()
	validate.RegisterValidation("pin", mustValidPin)

	type LoginRequest struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	loginRequest := LoginRequest{
		Phone: "01010101",
		Pin:   "123456",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err)
	}
}

func TestOrRule(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email|number"`
		Password string `validate:"required"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "x@x.x",
		Password: "Svelte",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err)
	}
}

func mustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, kind, isNullable, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Field is not ok")
	}

	fmt.Println("------------------------------")
	fmt.Println("value\t\t:", value)
	fmt.Println("kind\t\t:", kind)
	fmt.Println("isNullable\t:", isNullable)
	fmt.Println("ok\t\t:", ok)
	fmt.Println("------------------------------")

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomValidationCrossField(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", mustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Phone|field_equals_ignore_case=Email"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,number"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "x@x.x",
		Email:    "x@x.x",
		Phone:    "01010101",
		Name:     "Svelte",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func mustValidRegisterSucces(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// success
	} else {
		// fail
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(mustValidRegisterSucces, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "x",
		Email:    "x@x.x",
		Phone:    "01010101",
		Password: "12345678",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err)
	}
}
