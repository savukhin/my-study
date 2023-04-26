package api

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "pong"})
}

func ExecuteQuery(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "pong"})
}

// func Auth(c *gin.Context) {}
