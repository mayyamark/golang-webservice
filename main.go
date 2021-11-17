package main

import (
	"fmt"

	"github.com/mayyamark/golang-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Mayya",
		LastName:  "Markova",
	}

	fmt.Println(u)
}
