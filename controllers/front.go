package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// set up the root routing
func RegisterControllers() {
	uc := newUserController()

	// any /users & /users/[id] is handled by the userController
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// encode the go object into a json
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
