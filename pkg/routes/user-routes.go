package routes

import (
	"github.com/gorilla/mux"
	"github.com/Deepanshu-iiits-it/go-authentication/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router){
	router.HandleFunc("/signup/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/signin/", controllers.LoginUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}