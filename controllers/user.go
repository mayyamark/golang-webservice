package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mayyamark/golang-webservice/models"
)

// handle the routing of http requests comes into our web server to the correct method that's going to handle that in the model layer
type userController struct {
	userIDPattern *regexp.Regexp // // userIDPattern - to check which type of request it's going to work with (to match the incomming http req path)
}

// 1. the signature (uc userController) binds the func to the userController type => the function becomes a method
// 2. we are creating a ServeHTTP method with those 2 params (Req & Res objects from the net/http package)
// => for go this automatically means that this method is an implementation of the Handler interface https://golang.org/pkg/http/#Handler
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// recieve the http req in and decide based in the info in that req,
	// which method from the ones below, to pass that req off to and
	// to be actually process

	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// userIDPattern is a regex => there is a build-in FindStringSubmatch method,
		// which compares the incoming url path to the regex we specified
		// it will return a slice with all of the matches => in our case
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		// if no matches => 404
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		// matches[1] is the subgroup match, that contain the id value
		// converting the string to a number
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			// if no matches => 404
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// retrieve all users from the models layer and returning it back out
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

// return a single user
func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	encodeResponseAsJSON(u, w)
}

// add a new user
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// type conversion: byte to string
		w.Write([]byte("could not parse User object"))
		return
	}

	// use = => overwriting u and err variables
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

// update an existing user
func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// type conversion: byte to string
		w.Write([]byte("could not parse User object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

// delete an existing user
func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	// convert the passed user from the req body and to a json object
	dec := json.NewDecoder(r.Body)

	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// a constructor function - by convention: starts with new
// constructs a new userController object
// returns a pointer to the userController
func newUserController() *userController {
	// create a userController and return the address of it
	return &userController{
		// (\d+) => defines a subgroup
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
