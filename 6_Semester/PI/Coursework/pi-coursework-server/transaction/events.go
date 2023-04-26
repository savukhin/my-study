package transaction

import (
	"encoding/json"
)

type IEvent interface {
	GetDescription() string
	GetEventType() string
	GetTableName() string
}

type Event struct {
	IEvent
	TableName string
}

type CreateEvent struct {
	*Event  `json:"-"`
	Lines   [][]string `json:"lines"`
	Columns []string   `json:"columns"`
}

type DeleteEvent struct {
	*Event  `json:"-"`
	Indexes []int32 `json:"indexes"`
}

type EventType string

const (
	CreateEventType EventType = "create"
	DeletEventType  EventType = "delete"
)

func (event *Event) GetTableName() string {
	return event.TableName
}

func (event *CreateEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *DeleteEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *CreateEvent) GetEventType() string {
	return string(CreateEventType)
}

func (event *DeleteEvent) GetEventType() string {
	return string(DeletEventType)
}
