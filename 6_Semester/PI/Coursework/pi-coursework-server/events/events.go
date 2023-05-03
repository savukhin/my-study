package events

import "encoding/json"

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
	*Event    `json:"-"`
	Lines     [][]string `json:"lines"`
	Columns   []string   `json:"columns"`
	TableName string     `json:"table_name"`
}

type InsertEvent struct {
	*Event     `json:"-"`
	AfterIndex int32    `json:"after_index"`
	Line       []string `json:"line"`
}

type UpdateEvent struct {
	*Event  `json:"-"`
	Indexes []int32           `json:"index"`
	Values  map[string]string `json:"values"`
}

type DeleteEvent struct {
	*Event  `json:"-"`
	Indexes []int32 `json:"indexes"`
}

type RollbackEvent struct {
	IEvent
	RollbackTransactionName string `json:"rollback_transaction_name"`
}

type EventType string

const (
	CreateEventType   EventType = "create"
	InsertEventType   EventType = "create"
	UpdateEventType   EventType = "create"
	DeletEventType    EventType = "delete"
	RollbackEventType EventType = "create"
)

func (event *Event) GetTableName() string {
	return event.TableName
}

func (event *RollbackEvent) GetTableName() string {
	return "None"
}

func (event *CreateEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *InsertEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *UpdateEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *DeleteEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *RollbackEvent) GetDescription() string {
	result, _ := json.Marshal(event)
	return string(result)
}

func (event *CreateEvent) GetEventType() string {
	return string(CreateEventType)
}

func (event *DeleteEvent) GetEventType() string {
	return string(DeletEventType)
}

func (event *InsertEvent) GetEventType() string {
	return string(InsertEventType)
}

func (event *UpdateEvent) GetEventType() string {
	return string(UpdateEventType)
}

func (event *RollbackEvent) GetEventType() string {
	return string(RollbackEventType)
}
