package transaction

import (
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

func (trans *ComplexTransaction) Eval(storage table.Storage, transactionLog *TransactionFile) (table.Storage, error) {
	nextStorage := *storage.Copy()
	events := make([]events.IEvent, len(trans.Executors))

	for i, exec := range trans.Executors {
		tmpStorage, event, err := exec.DoExecute(&nextStorage)
		nextStorage = tmpStorage

		if err != nil {
			return table.Storage{}, err
		}

		// events[i] = exec.ToEvent()
		events[i] = event
	}

	if transactionLog != nil {
		transactionLog.addTransaction(events, "")
	}

	return nextStorage, nil
}
