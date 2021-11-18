package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp // to match the invomming http req path
}

// bind the function to the userController
func (uc userController) serverHTTP(w http.ResponseWriter, r *http.Request) {
	// type conversion: byte to string
	w.Write([]byte("Hello from the User Controller!"))
}

// a construction function
func newUserController() *userController {
	// create a userController and return the address of it
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
