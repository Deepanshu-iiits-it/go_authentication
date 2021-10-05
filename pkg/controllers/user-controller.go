package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Deepanshu-iiits-it/go-authentication/pkg/models"
	"github.com/Deepanshu-iiits-it/go-authentication/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var NewUser models.User

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	newUsers:=models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request){
	creds := &models.User{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	e:= creds.Email
	p:= creds.Password
	userDetails, str, _ :=models.LoginUser(e, p)
	if str == "No"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	res, _:=json.Marshal(userDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	Createuser:= &models.User{}
	// utils.ParseBody(r,Createuser)
	err := json.NewDecoder(r.Body).Decode(Createuser)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Createuser.Password), 8)
	
	Createuser.Password = string(hashedPassword)
	u:= Createuser.CreateUser()
	res, _:=json.Marshal(u)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err!=nil {
		fmt.Println("error while parsing id")
	}
	user := models.DeleteUser(ID)
	res, _:=json.Marshal(user)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId:= vars["id"]
	ID, err := strconv.ParseInt(userId, 0,0)
	if err!=nil {
		fmt.Println("error while parsing id")
	}
	userDetails, db:=models.GetUserById(ID)
	if updateUser.Name !=""{
		userDetails.Name= updateUser.Name
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Password != ""{
		userDetails.Password = updateUser.Password
	}
	db.Save(&userDetails)
	res, _:= json.Marshal(userDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}