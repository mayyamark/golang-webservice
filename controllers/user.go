package controllers

import (
	"net/http"
	"regexp"
)

// handle the routing of http requests comes into our web server to the correct method that's going to handle that in the model layer
type userController struct {
	userIDPattern *regexp.Regexp // // userIDPattern - to check which type of request it's going to work with (to match the incomming http req path)
}

// 1. the signature (uc userController) binds the func to the userController type => the function becomes a method
// 2. we are creating a ServeHTTP method with those 2 params (Req & Res objects from the net/http package)
// => for go this automatically means that this method is an implementation of the Handler interface https://golang.org/pkg/http/#Handler
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// type conversion: byte to string
	w.Write([]byte("Hello from the User Controller!"))
}

// a constructor function - by convention: starts with new
// constructs a new userController object
// returns a pointer to the userController
func newUserController() *userController {
	// create a userController and return the address of it
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
