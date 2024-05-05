package email

import (
	"fmt"
	"github.com/SebastianOraczek/internal/env"
	"github.com/SebastianOraczek/internal/response"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/mail"
	"strings"
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

func CreateNameAndSurnameFromEmail(email string) (name string, surname string) {
	separatorIndex := strings.Index(email, "@")
	personData := strings.Split(email[0:separatorIndex], ".")

	name = cases.Title(language.English).String(personData[0])
	surname = cases.Title(language.English).String(personData[1])

	return name, surname
}
