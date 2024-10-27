package routes

import (
	"MatchMaker/controllers"

	"github.com/gin-gonic/gin"
)

// these routes are protected through the auth module
func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUpUser)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.AuthMiddleware(), controllers.Logout)
}
