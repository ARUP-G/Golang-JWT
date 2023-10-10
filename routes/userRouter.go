package routes

import (
	"github.com/ARUP-G/Golang-JWT/controller"
	"github.com/ARUP-G/Golang-JWT/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {

	// To authenticate after user have some token
	incomingRoutes.Use(middleware.Authenticate())

	// Get all user
	incomingRoutes.GET("/users", controller.GetUsers())

	// Get user by id
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
