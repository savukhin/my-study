package main

import (
	"coursework/shared/models"
	"coursework/shared/utils"
	"fmt"
	"log"
	"net"

	"github.com/go-gin/gin"
	"google.golang.org/grpc"

	pb "coursework/shared/proto/golang/messages"
)

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
	port := utils.GetEnvInt("SERVER_PORT", 3000)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCourseWorkServer(grpcServer, NewServer())
	grpcServer.Serve(lis)
}
