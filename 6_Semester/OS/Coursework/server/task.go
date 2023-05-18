package main

import (
	pb "coursework/shared/proto/golang/messages"
)

var tasksCounter int32 = 0

type Task struct {
	TaskID         int32
	ProcessTimeSec int32
	NeedResponse   bool
	Completed      bool
	Priority       int32
}

func GenerateTaskID() int32 {
	tasksCounter += 1
	return tasksCounter
}

func NewTask(processTimeSec int32, needResponse bool) Task {
	return Task{
		TaskID:         GenerateTaskID(),
		ProcessTimeSec: processTimeSec,
		NeedResponse:   needResponse,
		Completed:      false,
		Priority:       0,
	}
}

func FromGeneratorRequest(req pb.GeneratorRequest) Task {
	return Task{
		TaskID:         GenerateTaskID(),
		ProcessTimeSec: req.ProcessTimeSec,
		NeedResponse:   req.NeedResponse,
		Completed:      false,
		Priority:       0,
	}
}

func (task Task) ToProto() *pb.Task {
	return &pb.Task{
		TaskId:         task.TaskID,
		ProcessTimeSec: task.ProcessTimeSec,
		NeedResponse:   task.NeedResponse,
		Priority:       task.Priority,
	}
}
