package routes

import (
	"MatchMaker/controllers"

	"github.com/gin-gonic/gin"
)

// these routes are protected through the auth module
func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
}
