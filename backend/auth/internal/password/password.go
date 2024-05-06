package password

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/response"
	"regexp"
)

func Validate(password, confirmPassword string) error {
	if len(password) < 8 {
		return fmt.Errorf(response.MinCharacterLengthErrMsg)
	}

	if !regexp.MustCompile("[A-Z]+").MatchString(password) {
		return fmt.Errorf(response.UppercaseCharacterErrMsg)
	}

	if !regexp.MustCompile("[0-9]+").MatchString(password) {
		return fmt.Errorf(response.MissingNumberErrMsg)
	}

	specialChars := "[!@#$%^&*()_+\\-=\\[\\]{}|\\\\,.?/<>]"
	if !regexp.MustCompile(specialChars).MatchString(password) {
		return fmt.Errorf(response.MissingSpecialCharacterErrMsg)
	}

	if confirmPassword != "" && password != confirmPassword {
		return fmt.Errorf(response.PasswordNotTheSameErrMsg)
	}

	return nil
}
