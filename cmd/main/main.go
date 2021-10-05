package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Deepanshu-iiits-it/go-authentication/pkg/routes"
)

func main() {
	r:= mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:3030", r))
}