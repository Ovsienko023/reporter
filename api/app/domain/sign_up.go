package domain

import (
	"errors"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type SignUpRequest struct {
	Login       string  `json:"login,omitempty"`
	Password    string  `json:"password,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

func (r *SignUpRequest) ToDbSignUp() *database.SignUp {
	return &database.SignUp{
		Login:       r.Login,
		Password:    r.Password,
		DisplayName: r.DisplayName,
	}
}

func (r SignUpRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Login, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Password, validation.Length(6, 128)),
		validation.Field(&r.Password, validation.By(checkPassword)),
		validation.Field(&r.DisplayName, validation.Required, validation.Length(0, 128)),
	)
}

func checkPassword(value interface{}) error {
	s, _ := value.(string)

	digit := regexp.MustCompile("[0-9]+")
	isDigit := digit.Match([]byte(s))

	lower := regexp.MustCompile("[a-z]+")
	isLower := lower.Match([]byte(s))

	upper := regexp.MustCompile("[A-Z]+")
	isUpper := upper.Match([]byte(s))

	if isDigit && isLower && isUpper {
		return nil
	}

	return errors.New("must contain numbers, latin letters of different case")
}
