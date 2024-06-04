// GOAL: Process and Validate HTML Form in GO Web Application

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	HOST = "127.0.0.1"
	PORT = ":8080"
)

func main() {
	router := setupRouter()

	log.Fatalln(router.Run(HOST + PORT))

}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// setup route handlers

	router.GET("/", Home)
	// router.GET("/")
	router.GET("/confirmation", Confirmation)

	return router

}

// Home: handles rendering the home page
func Home(c *gin.Context) {
	render(c.Writer, "http_servers_and_testing/contact_form/templates/home.html", nil)
}

func ValidateForm(c *gin.Context) {
	// Step 1: Validate form

	// Step 2: Send message in an email

	// Step 3: Redirect to confirmation page
}

func Confirmation(c *gin.Context) {
	render(c.Writer, "http_servers_and_testing/contact_form/templates/confirmation.html", nil)
}

// render: Renders the specified html document to the client.
func render(w http.ResponseWriter, filename string, data interface{}) {

	// create new template
	tmpl, err := template.ParseFiles(filename)

	if err != nil {
		log.Printf("render: error when attempting to render: %s | error: %s", filename, err.Error())
		http.Error(w, fmt.Sprintf("there was an error when attempting to read file: %s", filename), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("render: error when attempting to Execute teplate| error: %s", err.Error())
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

// TODO: Reword in your own words

// In essence:

//   - a Go template acts as a blueprint for generating HTML
//     where the actual content is dynamically filled in
//     based on data and logic from your application
