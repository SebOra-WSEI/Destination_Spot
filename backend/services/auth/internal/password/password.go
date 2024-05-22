package password

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func Validate(password, confirmPassword string) error {
	if len(password) < 8 {
		return response.ErrMinCharacterLength
	}

	if !regexp.MustCompile("[A-Z]+").MatchString(password) {
		return response.ErrUppercaseCharacter
	}

	if !regexp.MustCompile("[0-9]+").MatchString(password) {
		return response.ErrMissingNumber
	}

	specialChars := "[!@#$%^&*()_+\\-=\\[\\]{}|\\\\,.?/<>]"
	if !regexp.MustCompile(specialChars).MatchString(password) {
		return response.ErrMissingSpecialCharacter
	}

	if confirmPassword != "" && password != confirmPassword {
		return response.ErrPasswordNotTheSame
	}

	return nil
}

func Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		return "", response.ErrInternalServer
	}

	return string(hashedPassword), nil
}
