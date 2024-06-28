package main

import (
	"MatchMaker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	addRoutes(r)
	r.Run()
}

func addRoutes(r *gin.Engine) {
	routes.AuthRoutes(r)
}
