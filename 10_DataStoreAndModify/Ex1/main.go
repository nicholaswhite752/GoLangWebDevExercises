package main

import (
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/06_hands-on/starting-code/controllers"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}
