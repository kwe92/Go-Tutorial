// GOAL: Process and Validate HTML Form in GO Web Application

package main

import (
	"contact-form-validation/handler"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	HOST = "127.0.0.1"
	PORT = ":8080"
)

func main() {

	router := setupRouter()

	// start server
	log.Fatalln(router.Run(HOST + PORT))

}

// setupRouter: create routes and add middleware.
func setupRouter() *gin.Engine {
	router := gin.Default()

	// setup route handlers

	router.GET("/", handler.Home)
	router.POST("/", handler.ValidateForm)
	router.GET("/confirmation", handler.Confirmation)

	return router

}
