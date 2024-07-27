package controllers

import (
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
