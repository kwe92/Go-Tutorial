package handler

import (
	model "contact-form-validation/models"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home: handles rendering the home page
func Home(c *gin.Context) {
	// add cookie header data
	http.SetCookie(c.Writer, &http.Cookie{Name: "test", Value: "tcookie", Path: "/", Domain: "127.0.0.1:8080"})

	render(c.Writer, "http_servers_and_testing/contact_form/templates/home.html", nil)
}

func ValidateForm(c *gin.Context) {
	// Step 1: Retrieve form values from the request and validate the form
	contactForm := &model.ContactForm{
		Name:    c.Request.FormValue("name"),
		Email:   c.Request.FormValue("email"),
		Content: c.Request.FormValue("content"),
	}

	log.Printf("received contact form: %+v", contactForm)

	if !contactForm.Validate() {
		// re-render home page with updated errors if validation fails
		render(c.Writer, "http_servers_and_testing/contact_form/templates/home.html", contactForm)
		log.Printf("received contact form with errors: %+v", contactForm)

		return
	}

	// Step 2: Send contact form message in an email
	if err := contactForm.SendConfirmationEmail(); err != nil {
		log.Print(err)
		http.Error(c.Writer, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Println("c.Request.Header:", c.Request.Header)

	// Step 3: Redirect to confirmation page
	// if validation is successful redirect the user to a new page
	http.Redirect(c.Writer, c.Request, "/confirmation", http.StatusSeeOther)
}

func Confirmation(c *gin.Context) {
	render(c.Writer, "http_servers_and_testing/contact_form/templates/confirmation.html", nil)
}

// render: Renders the specified HTML document to the client dynamically changing template data if passed in.
func render(w http.ResponseWriter, filename string, data interface{}) {

	// create new template
	tmpl, err := template.ParseFiles(filename)

	if err != nil {
		log.Printf("render: error when attempting to render: %s | error: %s", filename, err.Error())
		http.Error(w, fmt.Sprintf("there was an error when attempting to read file: %s", filename), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("render: error when attempting to Execute teplate | error: %s", err.Error())
		http.Error(w, "something went wrong when attempting to render template. ", http.StatusInternalServerError)
		return
	}
}

// html/template

//   - a package used to safely render HTML documents to clients from a GO web server

//   - has built in features to protect against common web vulnerabilities like Cross-Site Scripting (XSS)

// When to Use The html/template Package to Render HTML

//   - whenever you need to render dynamic HTML documents to clients in GO

//   - when dealing with user generated data

// Why The html/template Package is Safer

//   - Automatic HTML special character exscaping protects against scripting injection due to XSS attacks

//   - data can be validated and sanitized before being processed into an HTML template

//   - granular control over dynamic HTML output

// Retrieving Form Values With Gin Context

//  - the Request object has a FormValue function

//  - use the FormValue to retrieve form values by key from the request

// Summary:

//   - Go templates act as blueprints for generating dynamic HTML

// TODO: look into https://pkg.go.dev/github.com/gorilla/schema for converting structs to and from forms
