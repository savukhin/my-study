package models

import "time"

type GeneratorRequest struct {
	ProcessTime  time.Duration `json:"process_time"`
	NeedResponse bool          `json:"need_response"`
	Priority     int           `json:"priority"`
}

type ProcessorMessage struct {
	ProcessTime time.Duration `json:"process_time"`
}

type ProcessorResponse struct {
	Status bool `json:"status"`
}

type GeneratorResponse struct {
	Status bool `json:"status"`
}
