package routes

import (
	"github.com/ARUP-G/Golang-JWT/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	// Singup function
	incomingRoutes.POST("user/singup", controller.Singup())

	// Login function
	incomingRoutes.POST("user/login", controller.Login())
}
