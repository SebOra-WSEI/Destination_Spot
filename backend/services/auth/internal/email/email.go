package email

import (
	 "github.com/SebOra-WSEI/Destination_spot/auth/internal/message"
	"github.com/SebOra-WSEI/Destination_spot/shared/env"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"net/mail"
	"strings"
)

func Validate(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return message.ErrInvalidEmailFormat
	}

	domain, err := env.GetEnvVariableByName(env.Domain)
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(email) < len(domain) || email[len(email)-len(domain):] != domain {
		return message.ErrInvalidDomain
	}

	return nil
}

func GetNameAndSurname(email string) (name string, surname string) {
	separatorIndex := strings.Index(email, "@")
	personData := strings.Split(email[0:separatorIndex], ".")

	name = cases.Title(language.English).String(personData[0])
	surname = cases.Title(language.English).String(personData[1])

	return name, surname
}
