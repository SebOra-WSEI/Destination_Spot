package email

import (
	"fmt"
	"github.com/SebOra-WSEI/auth/internal/env"
	"github.com/SebOra-WSEI/auth/internal/response"
	"net/mail"
)

func Verify(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf(response.InvalidEmailFormatErrorMsg)
	}

	domain := env.GetEnvVariableByName(env.Domain)
	if len(email) < len(domain) || email[len(email)-len(domain):] != domain {
		return fmt.Errorf(response.InvalidDomainErrorMsg)
	}

	return nil
}
