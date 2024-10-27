package controllers

import (
	"MatchMaker/auth"
	"MatchMaker/database"
	"MatchMaker/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /signup
func SignUpUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	_, err = database.DB.Exec(`INSERT INTO [User] (FirstName, LastName, Email, Password) VALUES (@p1, @p2, @p3, @p4)`,
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Printf("Fejl ved indsættelse af bruger i databasen: %v", err)
		return
	}

	// Return success response
	c.JSON(http.StatusOK, user)
}

// POST /login
func Login(c *gin.Context) {
	var user models.Login

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedPassword string
	err := database.DB.QueryRow(`SELECT Password FROM [User] WHERE Email = @p1`, user.Email).Scan(&storedPassword)

	// Hvis brugeren ikke findes
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Sammenlign adgangskoden med den hashede adgangskode i databasen
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	jwtToken, err := auth.CreateToken(user.Email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.SetCookie("Authorization", jwtToken, auth.ExpirationTimeSeconds, "/", "localhost", true, true)
	c.Status(http.StatusOK)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify the token (assuming you have a method in the auth package to validate the token)
		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set user information from the token in the request context
		c.Set("user", sub)

		// Proceed to the next middleware or route handler
		c.Next()
	}
}

// GET /logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}
