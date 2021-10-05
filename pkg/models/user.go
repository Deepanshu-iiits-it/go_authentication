package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Deepanshu-iiits-it/go-authentication/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type User struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(ID int64) (*User, *gorm.DB){
	var getUser User
	db:= db.Where("ID=?", ID).Find(&getUser)
	return &getUser, db
}

func LoginUser(e string, p string) (*User, string, *gorm.DB){
	var getUser User
	db := db.Where("Email=?", e).Find(&getUser)
	if err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(p)); err!=nil{
		return &getUser, "No", db
	}
	return &getUser, "Yes", db
}

func DeleteUser(ID int64) User{
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}

