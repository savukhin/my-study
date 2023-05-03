package transaction

import (
	"pi-coursework-server/events"
	"strconv"
	"time"
)

type TransactionFile struct {
	Logs []*Log
}

func NewTransactionFile() *TransactionFile {
	return &TransactionFile{}
}

func (logs *TransactionFile) addTransaction(events []*events.IEvent, transactionName string) {
	if transactionName == "" {
		transactionName = strconv.Itoa(len(logs.Logs))
	}
	transactionTime := time.Now()

	for _, event := range events {
		log := &Log{
			ev:              *event,
			TransactionTime: transactionTime,
			TransactionName: transactionName,
		}

		logs.Logs = append(logs.Logs, log)
	}

}

func (logs *TransactionFile) AddSingleEvent(event *events.IEvent, transactionName string) {
	evs := make([]*events.IEvent, 1)
	evs[0] = event
	logs.addTransaction(evs, transactionName)
}

func (logs *TransactionFile) AddCreateEvent(columns []string, lines [][]string, tableName, transactionName string) {
	abs := &events.Event{TableName: tableName}

	event := &events.CreateEvent{
		Lines:   lines,
		Columns: columns,
		Event:   abs,
	}

	event.Event.IEvent = event

	logs.AddSingleEvent(&event.IEvent, transactionName)

}

func (logs *TransactionFile) AddDeleteEvent(indexes []int32, tableName, transactionName string) {
	abs := &events.Event{TableName: tableName}

	event := &events.DeleteEvent{
		Indexes: indexes,
		Event:   abs,
	}
	event.Event.IEvent = event

	logs.AddSingleEvent(&event.IEvent, transactionName)
}
