package transaction

import (
	"fmt"
	"pi-coursework-server/events"
	"pi-coursework-server/executor"
	"pi-coursework-server/table"
)

type ComplexTransaction struct {
	Executors []executor.IExecutor
}

func NewComplexTransaction(executors []executor.IExecutor) *ComplexTransaction {
	return &ComplexTransaction{
		Executors: executors,
	}
}

func (trans *ComplexTransaction) Eval(storage table.Storage, transactionLog *TransactionFile, transactionName string, transactionTimeUnix int, save bool) (table.Storage, error) {
	nextStorage := *storage.Copy()
	events := make([]events.IEvent, len(trans.Executors))
	fmt.Println("read executors", trans.Executors)

	for i, exec := range trans.Executors {
		fmt.Println("executing", exec)
		tmpStorage, event, err := exec.DoExecute(&nextStorage)
		nextStorage = tmpStorage

		if err != nil {
			return table.Storage{}, err
		}

		// events[i] = exec.ToEvent()
		events[i] = event
		fmt.Println("event = ", event)
	}

	if transactionLog != nil {
		transactionLog.addTransaction(events, transactionName, transactionTimeUnix, save)
	}

	return nextStorage, nil
}

func ComplexFromEvents(eventsArr []events.IEvent) (*ComplexTransaction, error) {
	execs := make([]executor.IExecutor, len(eventsArr))

	for i, event := range eventsArr {
		ex, err := executor.FromEvent(event)
		if err != nil {
			return nil, err
		}

		fmt.Println("add exec", ex)
		execs[i] = ex
	}

	complex := NewComplexTransaction(execs)
	fmt.Println("executors", complex.Executors)
	return complex, nil
}
