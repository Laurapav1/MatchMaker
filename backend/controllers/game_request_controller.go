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
	fmt.Println(c.GetString("userEmail"))

	input.UserEmail = c.GetString("userEmail")

	_, err := database.DB.Exec("INSERT INTO GameRequest (UserEmail, Niveau, Location, Time, Gender, Amount, Price) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)",
		input.UserEmail, input.Niveau, input.Location, input.Time, input.Gender, input.Amount, input.Price)
	if err != nil {
		log.Fatal("Error inserting GameRequest entry: ", err.Error())
	}
	fmt.Println("GameRequest entry inserted successfully! before")

	fmt.Println(c.GetString("userEmail"))

	c.JSON(http.StatusCreated, gin.H{
		"message": "GameRequest entry inserted successfully last one!",
	})

}

// GET // Get a game{id}
func GetGameRequest(c *gin.Context) {
	rows, err := database.DB.Query("SELECT ID, Niveau, Location, Time, Gender, Amount, Price FROM GameRequest")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}
	defer rows.Close()

	var results []models.GameRequest
	for rows.Next() {
		var cs models.GameRequest
		if err := rows.Scan(&cs.ID, &cs.Niveau, &cs.Location, &cs.Time, &cs.Gender, &cs.Amount, &cs.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		results = append(results, cs)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Row iteration error"})
		return
	}

	// Send the results as JSON
	c.JSON(http.StatusOK, results)
}

// Delete a game/{id}
func DeleteGame(c *gin.Context) {
	gameID := c.Param("id")

	result, err := database.DB.Exec("DELETE FROM GameRequest WHERE ID = @p1", gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	// If no rows were affected, the game doesn't exist
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

// Put / Change a game/{id}
func ChangeGame(c *gin.Context) {
	gameID := c.Param("id")

	// Bind the request JSON to a GameRequest struct
	var input models.GameRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Execute the update statement in the database
	result, err := database.DB.Exec(`
		UPDATE GameRequest 
		SET Niveau = @p1, Location = @p2, Time = @p3, Gender = @p4, Amount = @p5, Price = @p6
		WHERE ID = @p7`,
		input.Niveau, input.Location, input.Time, input.Gender, input.Amount, input.Price, gameID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}

	// Check how many rows were affected (should be 1 if the game existed and was updated)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
		return
	}

	// If no rows were affected, the game doesn't exist
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}
