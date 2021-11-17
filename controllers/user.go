package controllers

import "net/http"

type userController struct{}

// bind the function to the userController
func (uc userController) serverHTTP(w http.ResponseWriter, r *http.Request) {
	// type conversion: byte to string
	w.Write([]byte("Hello from the User Controller!"))
}
