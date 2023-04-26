package transaction

import (
	"strconv"
	"time"
)

type TransactionFile struct {
	Logs []*Log
}

func NewTransactionFile() *TransactionFile {
	return &TransactionFile{}
}

func (logs *TransactionFile) addEvent(name string, event *Event) {
	if name == "" {
		name = strconv.Itoa(len(logs.Logs))
	}

	log := &Log{
		ev:              event.IEvent,
		TransactionTime: time.Now(),
		TransactionName: name,
	}

	logs.Logs = append(logs.Logs, log)

}

func (logs *TransactionFile) AddCreateEvent(columns []string, lines [][]string, tableName, transactionName string) {
	abs := &Event{TableName: tableName}

	event := &CreateEvent{
		Lines:   lines,
		Columns: columns,
		Event:   abs,
	}

	event.Event.IEvent = event

	logs.addEvent(transactionName, event.Event)

}

func (logs *TransactionFile) AddDeleteEvent(indexes []int32, tableName, transactionName string) {
	abs := &Event{TableName: tableName}

	event := &DeleteEvent{
		Indexes: indexes,
		Event:   abs,
	}
	event.Event.IEvent = event

	logs.addEvent(transactionName, event.Event)
}
