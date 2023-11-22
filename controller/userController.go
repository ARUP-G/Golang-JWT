package controller

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var validate = validator.New()

func HashPassword() {

}

func VarifyPassworf() {

}

func Singup() {

}

func Login() {

}
func GetUser() {

}
func GetUsers() {

}
