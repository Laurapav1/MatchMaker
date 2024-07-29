package routes

import (
	"MatchMaker/controllers"

	"github.com/gin-gonic/gin"
)

func GameRequestRoutes(router *gin.Engine) {
	router.POST("/gamerequest", controllers.CreateGameRequest)
	router.GET("/gamerequest", controllers.GetGameRequest)
}
