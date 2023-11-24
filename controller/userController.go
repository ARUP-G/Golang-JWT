package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ARUP-G/Golang-JWT/database"
	"github.com/ARUP-G/Golang-JWT/helpers"
	"github.com/ARUP-G/Golang-JWT/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var validate = validator.New()

func HashPassword() {

}

func VarifyPassworf() {

}

func Singup() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var c, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		// Created user
		var user models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// This validates all the data user gives with the struct user
		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Count -> of the users
		// Here count is used to check pre-existing data of the users, if the data already exists the count will be more than one
		count, err := userCollection.CountDocuments(c, bson.M{"email": user.Email})
		defer cancle()

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occoured whlile checking the Email"})
		}

		count, err := userCollection.CountDocuments(c, bson.M{"contact": user.Contact})
		defer cancle()

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occoured whlile checking the Contact"})
		}

		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": " Email or Contact no already exists"})
		}
	}
}

func Login() {

}
func GetUser() gin.HandlerFunc {
	// Admin can access the user data only
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")

		// To check the user is admin or not
		if err := helpers.MatchUserTypeToUserid(ctx, userID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var c, cancle = context.WithTimeout(context.Background(), 100*time.Second)

		// Defined user
		var user models.User
		// Funding the user and decoding json in bson for go
		err := userCollection.FindOne(c, bson.M{"user_id": userID}).Decode(&user)
		defer cancle()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}

}
func GetUsers() {

}
