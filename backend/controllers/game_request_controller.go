package controllers

import (
	"MatchMaker/database"
	"MatchMaker/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


// POST // add a game
func CreateGameRequest(c *gin.Context) {
	var input models.GameRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO GameRequest (Niveau, Location, Time, Gender, Amount, Price) VALUES (@p1, @p2, @p3, @p4, @p5, @p6)",
		input.Niveau, input.Location, input.Time, input.Gender, input.Amount, input.Price)
	if err != nil {
		log.Fatal("Error inserting GameRequest entry: ", err.Error())
	}
	fmt.Println("GameRequest entry inserted successfully!")

	c.JSON(http.StatusCreated, gin.H{
		"message": "GameRequest entry inserted successfully!",
	})
}

// GET // Get a game{id}
func GetGameRequest(c *gin.Context) {
	getGameRequest, err := database.GetAllGameRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getGameRequest)
}


// GET all games
func GetAllGameRequests(c *gin.Context){

}

// Delete a game/{id}
func DeleteGame(c *gin.Context){

}

// Put / Change a game/{id}
func ChangeGame(c *gin.Context){

}
