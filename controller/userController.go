package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ARUP-G/Golang-JWT/database"
	"github.com/ARUP-G/Golang-JWT/helpers"
	"github.com/ARUP-G/Golang-JWT/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var validate = validator.New()

func HashPassword() {

}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect")
		check = false
	}
	return check, msg
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

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()

		user.UserID = user.ID.Hex()

		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UserType, *&user.UserID)

		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(c, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user was not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancle()
		ctx.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var c, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User // for the user of given Email

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Find the given Email from the collection
		err := userCollection.FindOne(c, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancle()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		// Password varification
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancle()
	}

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
