package controllers

import (
	"MatchMaker/database"
	"MatchMaker/models"
	"net/http"
	"time"

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
func GetAllGameRequests(c *gin.Context) {
	var results []models.GameRequest

	// Retrieve all game requests using GORM
	if err := database.DB.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	c.JSON(http.StatusOK, results)
}

type GameRequestQuery struct {
	User     string `form:"user"`
	MinLevel int    `form:"minLevel"`
	MaxLevel int    `form:"maxLevel"`
	Location string `form:"location"`
	From     string `form:"from"`
	To       string `form:"to"`
	Gender   string `form:"gender"`
}

func GetGameRequests(c *gin.Context) {
	var query GameRequestQuery
	// Bind query parameters to the struct
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	var gameRequests []models.GameRequest
	dbQuery := database.DB.Model(&models.GameRequest{})

	// Apply filters dynamically
	if query.User != "" {
		dbQuery = dbQuery.Where("user_email = ?", query.User)
	}
	if query.MinLevel > 0 {
		dbQuery = dbQuery.Where("niveau >= ?", query.MinLevel)
	}
	if query.MaxLevel > 0 {
		dbQuery = dbQuery.Where("niveau <= ?", query.MaxLevel)
	}
	if query.Location != "" {
		dbQuery = dbQuery.Where("location = ?", query.Location)
	}
	if query.Gender != "" {
		dbQuery = dbQuery.Where("gender = ?", query.Gender)
	}

	// Handle time range filters
	if query.From != "" {
		fromTime, err := time.Parse("2006-01-02T15:04", query.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' format. Use YYYY-MM-DDTHH:MM"})
			return
		}
		dbQuery = dbQuery.Where("time >= ?", fromTime)
	}
	if query.To != "" {
		toTime, err := time.Parse("2006-01-02T15:04", query.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'to' format. Use YYYY-MM-DDTHH:MM"})
			return
		}
		dbQuery = dbQuery.Where("time <= ?", toTime)
	}

	// Execute the query
	if err := dbQuery.Find(&gameRequests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve game requests"})
		return
	}

	c.JSON(http.StatusOK, gameRequests)

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
