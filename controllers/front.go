package controllers

import "net/http"

// set up the root routing
func RegisterControllers() {
	uc := newUserController()

	// any /users & /users/[id] is handled by the userController
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
