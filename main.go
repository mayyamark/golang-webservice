package main

import (
	"net/http"

	"github.com/mayyamark/golang-webservice/controllers"
)

func main() {
	// register our routing
	controllers.RegisterControllers()

	// listen on port 3000 and decide which route will handle the request
	http.ListenAndServe(":3000", nil)
}

// to start the app:
// 1. in the Terminal: go build .
// 2. in the Terminal: ./golang-webservice
// 3. in the browser: localhost:3000/users
