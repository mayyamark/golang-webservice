package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp // to match the invomming http req path
}

// bind the function to the userController
// we are creating a ServeHTTP method with those 2 params
// => for go this automatically means that this method is an implementation of the Handler type https://golang.org/pkg/http/#Handler
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
