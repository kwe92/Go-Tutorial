package model

import (
	"regexp"
	"strings"

	"github.com/go-mail/mail"
)

// regular expression to check if the input string is an email.
var rxEmail = regexp.MustCompile(".+@.+\\..+")

// ContactForm represents our applications contact form
type ContactForm struct {
	Name    string
	Email   string
	Content string
	Errors  map[string]string
}

// Validate: validates form if no errors are encountered the form is considered valid.
func (c *ContactForm) Validate() bool {
	// create a map to add errors to for each input
	c.Errors = make(map[string]string)

	if strings.TrimSpace(c.Name) == "" {
		c.Errors["name"] = "name can not be empty"
	}

	// check email validity
	validEmail := rxEmail.Match([]byte(c.Email))

	if !validEmail {
		c.Errors["email"] = "invalid email address"
	}

	if strings.TrimSpace(c.Content) == "" {
		c.Errors["content"] = "message can not be empty"
	}

	return len(c.Errors) == 0
}

// SendConfirmationEmail: Send comfirmation email to user upon successful form validation
func (c *ContactForm) SendConfirmationEmail() error {

	email := mail.NewMessage()

	email.SetHeader("To", "admin@example.com")

	email.SetHeader("From", "server@example.com")

	email.SetHeader("Reply-To", c.Email)

	email.SetHeader("Subject", "Contact form confirmation email")

	// credentails can be found in mailtrap sandbox

	username := "" //TODO: place your mailtrap username here

	password := "" //TODO: place your mailtrap password here

	return mail.NewDialer("sandbox.smtp.mailtrap.io", 25, username, password).DialAndSend(email)

}
