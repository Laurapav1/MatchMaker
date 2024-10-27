package main

import (
	"MatchMaker/database"
	"MatchMaker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	addRoutes(r)
	r.Run()
}

func addRoutes(r *gin.Engine) {
	routes.AuthRoutes(r)
	routes.GameRequestRoutes(r)
}
