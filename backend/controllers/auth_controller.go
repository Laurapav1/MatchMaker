package controllers

import (
	"MatchMaker/auth"
	"MatchMaker/database"
	"MatchMaker/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// POST /signup
func SignUpUser(c *gin.Context) {
	var user models.User

	// Bind JSON input to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create the user using GORM
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Printf("Error inserting user into the database: %v", err)
		return
	}

	// Return success response
	c.JSON(http.StatusOK, user)
}

// POST /login
func Login(c *gin.Context) {
	var loginData models.Login
	var user models.User

	// Bind JSON input to login struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user by email
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		} else {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a JWT token
	jwtToken, err := auth.CreateToken(user.Email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	// Set the JWT token as a cookie
	c.SetCookie("Authorization", jwtToken, auth.ExpirationTimeSeconds, "/", "localhost", true, true)
	c.Status(http.StatusOK)
}

// Middleware to protect routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify the token
		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract the subject (user email) from the token claims
		sub, err := token.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set user information from the token in the request context
		c.Set("userEmail", sub)

		// Proceed to the next middleware or route handler
		c.Next()
	}
}

// GET /logout
func Logout(c *gin.Context) {
	// Clear the Authorization cookie
	c.SetCookie("Authorization", "", -1, "/", "localhost", true, true)

	// Return a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
