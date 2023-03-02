package models

type GeneratorRequest struct {
	ProcessTimeSec int  `json:"process_time_sec"`
	NeedResponse   bool `json:"need_response"`
	Priority       int  `json:"priority"`
}

type GeneratorResponse struct {
	TaskID int `json:"task_id"`
}

type ProcessorMessage struct {
	ProcessTimeSec int `json:"process_time_sec"`
	TaskID         int `json:"task_id"`
}

type ProcessorResponse struct {
	Status bool `json:"status"`
}

type TaskCompletedRequest struct {
	Status bool `json:"status"`
	TaskID int  `json:"task_id"`
}
