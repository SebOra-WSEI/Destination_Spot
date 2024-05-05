package password

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/response"
	"regexp"
)

func Validate(password string) error {
	if len(password) < 8 {
		return fmt.Errorf(response.MinCharacterLengthErrorMsg)
	}

	if !regexp.MustCompile("[A-Z]+").MatchString(password) {
		return fmt.Errorf(response.UppercaseCharacterErrorMsg)
	}

	if !regexp.MustCompile("[0-9]+").MatchString(password) {
		return fmt.Errorf(response.MissingNumberErrorMsg)
	}

	specialChars := "[!@#$%^&*()_+\\-=\\[\\]{}|\\\\,.?/<>]"
	if !regexp.MustCompile(specialChars).MatchString(password) {
		return fmt.Errorf(response.MissingSpecialCharacterErrMsg)
	}

	return nil
}
