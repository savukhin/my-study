package transaction

import (
	"fmt"
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

func (logs *TransactionFile) addTransaction(evs []events.IEvent, transactionName string, transactionTimeUnix int, save bool) {
	if transactionName == "" {
		transactionName = strconv.Itoa(len(logs.Logs))
	}

	transactionTime := time.Now()

	if transactionTimeUnix != -1 {
		transactionTime = time.Unix(0, int64(transactionTimeUnix))
	}

	for _, event := range evs {
		log := &Log{
			Ev:              event,
			TransactionTime: transactionTime,
			TransactionName: transactionName,
		}

		logs.Logs = append(logs.Logs, log)
	}

	if save {
		logs.Save()
		fmt.Println("Saved")
	}

	if len(evs) == 1 {
		_, ok := evs[0].(*events.RollbackEvent)
		if ok {
			_, err := logs.ActiveComplexTransactions.Pop()
			fmt.Println("rollback", err)
			return
		}

		_, ok = evs[0].(*events.WriteEvent)
		if ok {
			fmt.Println("write")
			return
		}
		fmt.Println("nor")
	}

	logs.ActiveComplexTransactions.Push(transactionName)
}

func (logs *TransactionFile) AddSingleEvent(event events.IEvent, transactionName string, transactionTimeUnix int, save bool) {
	evs := make([]events.IEvent, 1)
	evs[0] = event
	logs.addTransaction(evs, transactionName, transactionTimeUnix, save)
}

func (logs *TransactionFile) AddCreateEvent(columns []string, lines [][]string, tableName, transactionName string, transactionTimeUnix int, save bool) {
	abs := &events.Event{}
	// abs := &events.Event{TableName: tableName}

	event := &events.CreateEvent{
		// Lines:   lines,
		// Columns: columns,
		TableName: tableName,
		Event:     abs,
	}

	event.Event.IEvent = event

	logs.AddSingleEvent(event.IEvent, transactionName, transactionTimeUnix, save)

}

func (logs *TransactionFile) AddDeleteEvent(indexes []int, tableName, transactionName string, transactionTimeUnix int, save bool) {
	abs := &events.Event{}
	// abs := &events.Event{TableName: tableName}

	event := &events.DeleteEvent{
		Indexes:   indexes,
		TableName: tableName,
		Event:     abs,
	}
	event.Event.IEvent = event

	logs.AddSingleEvent(event.IEvent, transactionName, transactionTimeUnix, save)
}

func (logs *TransactionFile) GetRollbackedComplexTransactionByName(name string) (ComplexTransaction, error) {
	executors := make([]executor.IExecutor, 0)
	// fmt.Println("Logs = ", logs.Logs)
	for _, log := range logs.Logs {
		fmt.Println("log rollback ev = ", log.Ev)
		if log.TransactionName != name {
			continue
		}

		_, writeOk := log.Ev.(*events.WriteEvent)
		if writeOk {
			continue
		}

		exec, err := executor.RollbackEvent(log.Ev)
		if err != nil {
			return ComplexTransaction{}, err
		}

		executors = append(executors, exec)
	}

	trans := NewComplexTransaction(utils.Reverse(executors))
	return *trans, nil
}

func (logs *TransactionFile) GetComplexTransactionByName(name string) (ComplexTransaction, error) {
	executors := make([]executor.IExecutor, 0)
	for _, log := range logs.Logs {
		if log.TransactionName != name {
			continue
		}

		exec, err := executor.FromEvent(log.Ev)
		if err != nil {
			return ComplexTransaction{}, err
		}

		executors = append(executors, exec)
	}

	trans := NewComplexTransaction(utils.Reverse(executors))
	return *trans, nil
}

func (logs *TransactionFile) GetLastActiveComplexTransactionName() (string, error) {
	return logs.ActiveComplexTransactions.Top()
}

func (logs *TransactionFile) GetLastWriteIndex() int {
	for i := len(logs.Logs) - 1; i >= 0; i-- {
		_, ok := logs.Logs[i].Ev.(*events.WriteEvent)
		if ok {
			return i
		}
	}
	return -1
}
