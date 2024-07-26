package controllers

import (
	"MatchMaker/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /signup
func Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Signup",
	})
}

// POST /login
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

// GET /logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}

type CreateSearchInput struct {
    Niveau   int     `json:"niveau"`
    Location string  `json:"location"`
    Time     string  `json:"time"`
    Gender   string  `json:"gender"`
    Amount   int     `json:"amount"`
    Price    float64 `json:"price"`
}

func CreateSearch(c *gin.Context) {
	var input CreateSearchInput
 
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO Search (Niveau, Location, Time, Gender, Amount, Price) VALUES (@p1, @p2, @p3, @p4, @p5, @p6)",
		input.Niveau, input.Location, input.Time, input.Gender, input.Amount, input.Price)
	if err != nil {
		log.Fatal("Error inserting createSearch entry: ", err.Error())
	}
	fmt.Println("Search entry inserted successfully!")

	c.JSON(http.StatusCreated, gin.H{
		"message": "Search entry inserted successfully!",
	})
}

//niveau -min og max
//sted
//tidspunkt
//køn
//antal
//pris
