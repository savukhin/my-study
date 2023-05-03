package transaction

import (
	"pi-coursework-server/events"
	"pi-coursework-server/executor"
	"pi-coursework-server/utils"
	"strconv"
	"time"
)

type TransactionFile struct {
	Logs                      []*Log
	ActiveComplexTransactions utils.Stack[string]
}

func NewTransactionFile() *TransactionFile {
	return &TransactionFile{}
}

func (logs *TransactionFile) addTransaction(evs []events.IEvent, transactionName string) {
	if transactionName == "" {
		transactionName = strconv.Itoa(len(logs.Logs))
	}
	transactionTime := time.Now()

	for _, event := range evs {
		log := &Log{
			ev:              event,
			TransactionTime: transactionTime,
			TransactionName: transactionName,
		}

		logs.Logs = append(logs.Logs, log)
	}

	if len(evs) == 1 {
		_, ok := evs[0].(*events.RollbackEvent)
		if ok {
			logs.ActiveComplexTransactions.Pop()
			return
		}
	}

	logs.ActiveComplexTransactions.Push(transactionName)
}

func (logs *TransactionFile) AddSingleEvent(event events.IEvent, transactionName string) {
	evs := make([]events.IEvent, 1)
	evs[0] = event
	logs.addTransaction(evs, transactionName)
}

func (logs *TransactionFile) AddCreateEvent(columns []string, lines [][]string, tableName, transactionName string) {
	abs := &events.Event{TableName: tableName}

	event := &events.CreateEvent{
		// Lines:   lines,
		// Columns: columns,
		Event: abs,
	}

	event.Event.IEvent = event

	logs.AddSingleEvent(event.IEvent, transactionName)

}

func (logs *TransactionFile) AddDeleteEvent(indexes []int, tableName, transactionName string) {
	abs := &events.Event{TableName: tableName}

	event := &events.DeleteEvent{
		Indexes: indexes,
		Event:   abs,
	}
	event.Event.IEvent = event

	logs.AddSingleEvent(event.IEvent, transactionName)
}

func (logs *TransactionFile) GetRollbackedComplexTransactionByName(name string) (ComplexTransaction, error) {
	executors := make([]executor.IExecutor, 0)
	for _, log := range logs.Logs {
		if log.TransactionName != name {
			continue
		}

		exec, err := executor.RollbackEvent(log.ev)
		if err != nil {
			return ComplexTransaction{}, err
		}

		executors = append(executors, exec)
	}

	trans := NewComplexTransaction(utils.Reverse(executors))
	return *trans, nil
}

func (logs *TransactionFile) GetLastActiveComplexTransactionName() string {
	return logs.ActiveComplexTransactions.Top()
}
