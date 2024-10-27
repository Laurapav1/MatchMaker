package controllers

import (
	"MatchMaker/database"
	"MatchMaker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST // add a game
// POST /game
func CreateGameRequest(c *gin.Context) {
	var input models.GameRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the user's email from the context
	input.UserEmail = c.GetString("userEmail")

	// Insert the new game request using GORM
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "GameRequest entry inserted successfully!",
	})
}

// GET /game
func GetGameRequest(c *gin.Context) {
	var results []models.GameRequest

	// Retrieve all game requests using GORM
	if err := database.DB.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// DELETE /game/{id}
func DeleteGame(c *gin.Context) {
	gameID := c.Param("id")

	// Delete the game using GORM
	if err := database.DB.Delete(&models.GameRequest{}, gameID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

// PUT /game/{id}
func ChangeGame(c *gin.Context) {
	gameID := c.Param("id")
	var input models.GameRequest

	// Bind the request JSON to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the game by ID and update it using GORM
	if err := database.DB.Model(&models.GameRequest{}).Where("id = ?", gameID).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}
