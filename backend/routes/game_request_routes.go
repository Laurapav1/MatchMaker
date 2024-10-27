package routes

import (
	"MatchMaker/controllers"

	"github.com/gin-gonic/gin"
)

func GameRequestRoutes(router *gin.Engine) {
	router.POST("/gamerequests", controllers.AuthMiddleware(), controllers.CreateGameRequest) //Add a Game
	router.GET("/gamerequests", controllers.AuthMiddleware(), controllers.GetGameRequest) // Get a game
	
}
