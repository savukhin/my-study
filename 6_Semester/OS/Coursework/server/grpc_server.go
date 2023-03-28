package main

import (
	"context"
	pb "coursework/shared/proto/golang/messages"
)

type CourseWorkServer struct {
	pb.UnimplementedCourseWorkServer
}

// type Hub struct {

// }

var (
	processorUrls = make([]string, 0)
	queue         PriorityQueue[Task]
	processors    = make(map[int]bool)
	taskIDToChan  = make(map[int32]chan Task)
)

func NewServer() CourseWorkServer {
	return CourseWorkServer{}
}

func CreateNewTasksChan() chan Task {
	tasksChan := make(chan Task)

	return tasksChan
}

func (serv CourseWorkServer) GenerateTasks(stream pb.CourseWork_GenerateTasksServer) error {
	tasksChan := make(chan Task)

	go func() {
		for task := range tasksChan {
			stream.Send(task.ToProto())
		}
	}()

	for {
		req, err := stream.Recv()
		if err != nil {
			continue
		}

		task := FromGeneratorRequest(*req)

		taskIDToChan[task.TaskID] = tasksChan

		queue.AddElement(req.Priority, task)

		stream.Send(task.ToProto())
	}

	return nil
}

func (serv CourseWorkServer) SubscribeProcessorOnTasks(_ *pb.Empty, stream pb.CourseWork_SubscribeProcessorOnTasksServer) error {

	return nil
}

func (serv CourseWorkServer) CompletedTask(ctx context.Context, taskPb *pb.CompletedTaskRequest) (*pb.Empty, error) {
	// taskIDToChan[taskPb.TaskId] <-
	return nil, nil
}
