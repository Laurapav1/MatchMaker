package main

import (
	"MatchMaker/database"
	"MatchMaker/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	r := gin.Default()
	addRoutes(r)
	r.Run()
}

func addRoutes(r *gin.Engine) {
	routes.AuthRoutes(r)
}
