package api

import (
	mainexecutor "pi-coursework-server/main_executor"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"message": "pong"})
}

func formResponseStr(responses []string) string {
	result := ""
	for _, output := range responses {
		if result == "" {
			result = output
			continue
		}
		if output == "" {
			continue
		}
		result = result + "\n\n" + output
	}

	return result
}

func ExecuteQuery(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// c.JSON(200, gin.H{"message": "pong"})
	query := &QeuryDTO{}
	if err := c.ShouldBindJSON(query); err != nil {
		c.JSON(422, gin.H{"message": "bad request"})
		return
	}

	responses, err := mainexecutor.ExecuteWholeQuery(query.Query)

	if err != nil {
		responses = append(responses, err.Error())
		c.JSON(400, gin.H{"message": formResponseStr(responses)})
		return
	}

	c.JSON(200, gin.H{"message": formResponseStr(responses)})
}

// func Auth(c *gin.Context) {}
