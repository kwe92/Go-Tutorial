package model

import (
	"regexp"
	"strings"

	"github.com/go-mail/mail"
)

// TODO: Review

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type ContactForm struct {
	Name    string
	Email   string
	Content string
	Errors  map[string]string
}

func (c *ContactForm) Validate() bool {
	c.Errors = make(map[string]string)

	if strings.TrimSpace(c.Name) == "" {
		c.Errors["name"] = "name can not be empty"
	}

	validEmail := rxEmail.Match([]byte(c.Email))

	if !validEmail {
		c.Errors["email"] = "invalid email address"
	}

	if strings.TrimSpace(c.Content) == "" {
		c.Errors["content"] = "message can not be empty"
	}

	return len(c.Errors) == 0
}

func (c *ContactForm) SendConfirmationEmail() error {

	email := mail.NewMessage()

	email.SetHeader("To", "admin@example.com")

	email.SetHeader("From", "server@example.com")

	email.SetHeader("Reply-To", c.Email)

	email.SetHeader("Subject", "Contact form confirmation email")

	// credentails can be found in mailtrap sandbox

	username := ""

	password := ""

	return mail.NewDialer("sandbox.smtp.mailtrap.io", 25, username, password).DialAndSend(email)

}
