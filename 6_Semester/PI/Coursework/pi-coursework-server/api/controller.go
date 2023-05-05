package api

import (
	"fmt"
	mainexecutor "pi-coursework-server/main_executor"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "pong"})
}

func ExecuteQuery(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// c.JSON(200, gin.H{"message": "pong"})
	query := &QeuryDTO{}
	if err := c.ShouldBindJSON(query); err != nil {
		c.JSON(422, gin.H{"message": "bad request"})
		return
	}

	response, err := mainexecutor.ExecuteWholeQuery(query.Query)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": response})
		return
	}

	c.JSON(200, gin.H{"message": response})
}

// func Auth(c *gin.Context) {}
