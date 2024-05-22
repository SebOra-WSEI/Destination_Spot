package password

import (
	"fmt"
	 "github.com/SebOra-WSEI/Destination_spot/auth/internal/message"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func Validate(password, confirmPassword string) error {
	if len(password) < 8 {
		return message.ErrMinCharacterLength
	}

	if !regexp.MustCompile("[A-Z]+").MatchString(password) {
		return message.ErrUppercaseCharacter
	}

	if !regexp.MustCompile("[0-9]+").MatchString(password) {
		return message.ErrMissingNumber
	}

	specialChars := "[!@#$%^&*()_+\\-=\\[\\]{}|\\\\,.?/<>]"
	if !regexp.MustCompile(specialChars).MatchString(password) {
		return message.ErrMissingSpecialCharacter
	}

	if confirmPassword != "" && password != confirmPassword {
		return message.ErrPasswordNotTheSame
	}

	return nil
}

func Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("Problem with hashing password", err.Error())
		return "", message.ErrInternalServer
	}

	return string(hashedPassword), nil
}
