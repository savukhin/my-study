package main

import (
	"coursework/shared/models"
	"fmt"

	"github.com/go-gin/gin"
)

type Task struct {
	TaskID         int
	ProcessTimeSec int
	NeedResponse   bool
}

var (
	processorUrls = make([]string, 0)
	queue         PriorityQueue[Task]
	tasksCounter  = 0
)

func GenerateTaskID() int {
	tasksCounter += 1
	return tasksCounter
}

func GeneratorRequest(c *gin.Context) {
	requestDto := &models.GeneratorRequest{}

	if err := c.BindJSON(&requestDto); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	task := Task{
		TaskID:         GenerateTaskID(),
		ProcessTimeSec: requestDto.ProcessTimeSec,
		NeedResponse:   requestDto.NeedResponse,
	}

	queue.AddElement(requestDto.Priority, task)

	c.JSON(200, models.GeneratorResponse{
		TaskID: task.TaskID,
	})
}

func PopTaskRequest(c *gin.Context) {
	_, task, err := queue.Pop()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "no tasks",
		})
	}

	dto := models.ProcessorMessage{
		TaskID:         task.TaskID,
		ProcessTimeSec: task.ProcessTimeSec,
	}

	c.JSON(200, dto)
}

func CompletedTaskRequest(c *gin.Context) {

}

func main() {
	r := gin.Default()
	fmt.Println(r)
	r.POST("/generator", GeneratorRequest)
	r.PATCH("/processor/pop-task", PopTaskRequest)
	r.PATCH("/processor/completed-task", CompletedTaskRequest)
}
