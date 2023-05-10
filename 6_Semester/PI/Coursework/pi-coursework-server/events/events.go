package events

import "encoding/json"

type IEvent interface {
	GetDescription() string
	GetEventType() string
	GetTableName() string
	Apply() error
}

type Event struct {
	IEvent
	TableName string
}

type CreateEvent struct {
	*Event    `json:"-"`
	TableName string   `json:"table_name"`
	Columns   []string `json:"columns"`
}

type InsertEvent struct {
	*Event `json:"-"`
	Values []string `json:"values"`
	// Values    map[string]string `json:"values"`
	Index     int    `json:"index"`
	TableName string `json:"table_name"`
}

type UpdateEvent struct {
	*Event    `json:"-"`
	Indexes   []int            `json:"index"`
	Values    map[int][]string `json:"values"`
	OldValues map[int][]string `json:"old_values"`
	TableName string           `json:"table_name"`
}

type DeleteEvent struct {
	*Event        `json:"-"`
	Indexes       []int            `json:"indexes"`
	DeletedValues map[int][]string `json:"deleted_values"`
	TableName     string           `json:"table_name"`
}

type DropEvent struct {
	*Event    `json:"-"`
	TableName string     `json:"table_name"`
	Columns   []string   `json:"columns"`
	Values    [][]string `json:"values"`
}

type RollbackEvent struct {
	// IEvent          `json:"-"`
	*Event          `json:"-"`
	TransactionName string `json:"rollback_transaction_name"`
}

type WriteEvent struct {
	// IEvent `json:"-"`
	*Event `json:"-"`
}

type EventType string

const (
	CreateEventType   EventType = "create"
	InsertEventType   EventType = "insert"
	UpdateEventType   EventType = "update"
	DeleteEventType   EventType = "delete"
	DropEventType     EventType = "drop"
	RollbackEventType EventType = "rollback"
	WriteEventType    EventType = "write"
)

func NewCreateEvent(tableName string, columns []string) *CreateEvent {
	abs := &Event{
		TableName: tableName,
	}
	event := &CreateEvent{
		TableName: tableName,
		Columns:   columns,
		Event:     abs,
	}
	abs.IEvent = event
	return event
}

func NewDropEvent(tableName string, columns []string, values [][]string) *DropEvent {
	abs := &Event{
		TableName: tableName,
	}
	event := &DropEvent{
		TableName: tableName,
		Columns:   columns,
		Values:    values,
		Event:     abs,
	}
	abs.IEvent = event
	return event
}

func NewInsertEvent(tableName string, values []string, index int) *InsertEvent {
	abs := &Event{
		TableName: tableName,
	}
	event := &InsertEvent{
		TableName: tableName,
		Values:    values,
		Index:     index,
		Event:     abs,
	}
	abs.IEvent = event
	return event
}

func NewDeleteEvent(tableName string, indexes []int, deletedValues map[int][]string) *DeleteEvent {
	abs := &Event{
		TableName: tableName,
	}
	event := &DeleteEvent{
		TableName:     tableName,
		Indexes:       indexes,
		DeletedValues: deletedValues,
		Event:         abs,
	}
	abs.IEvent = event
	return event
}

func NewUpdateEvent(tableName string, indexes []int, values map[int][]string, oldValues map[int][]string) *UpdateEvent {
	abs := &Event{
		TableName: tableName,
	}
	event := &UpdateEvent{
		TableName: tableName,
		Indexes:   indexes,
		Values:    values,
		OldValues: oldValues,
		Event:     abs,
	}
	abs.IEvent = event
	return event
}

func NewRollbackEvent(transactionName string) *RollbackEvent {
	abs := &Event{
		TableName: "None",
	}
	event := &RollbackEvent{
		TransactionName: transactionName,
		Event:           abs,
	}
	abs.IEvent = event

	return event
}

func NewWriteEvent() *WriteEvent {
	abs := &Event{
		TableName: "None",
	}
	event := &WriteEvent{}

	abs.IEvent = event
	return event
}

func (event *Event) GetTableName() string {
	return event.TableName
}

func (event *RollbackEvent) GetTableName() string {
	return "None"
}

func (event *WriteEvent) GetTableName() string {
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

func (event *WriteEvent) GetDescription() string {
	return "{}"
}

func (event *CreateEvent) GetEventType() string {
	return string(CreateEventType)
}

func (event *DeleteEvent) GetEventType() string {
	return string(DeleteEventType)
}

func (event *DropEvent) GetEventType() string {
	return string(DropEventType)
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

func (event *WriteEvent) GetEventType() string {
	return string(WriteEventType)
}
