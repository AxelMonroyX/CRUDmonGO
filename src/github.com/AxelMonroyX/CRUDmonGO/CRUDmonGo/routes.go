package main

import (
	// Standard library packages
	"net/http"

	"github.com/julienschmidt/httprouter"
	"./controllers"
	"gopkg.in/mgo.v2"
	"fmt"
)

func main() {
	// Instantiate a new router
	route := httprouter.New()

	// Get a user resource
	route.GET("/", Index)

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	route.GET("/user/:id", uc.GetUser)

	route.GET("/user", uc.GetAllUsers)

	// Create a new user
	route.POST("/user", uc.CreateUser)

	// Remove an existing user
	route.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", route)
}
// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome!</h1>\n")
	fmt.Fprint(w, "HOLAAAAA")
}