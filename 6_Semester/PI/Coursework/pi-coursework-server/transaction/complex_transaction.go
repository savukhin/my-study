package transaction

import (
	"pi-coursework-server/events"
	"pi-coursework-server/executor"
	"pi-coursework-server/table"
)

type ComplexTransaction struct {
	ITransaction

	Executors []executor.IExecutor
}

func NewComplexTransaction(executors []executor.IExecutor) *ComplexTransaction {
	return &ComplexTransaction{
		Executors: executors,
	}
}

func (trans *ComplexTransaction) Eval(storage table.Storage, transactionLog *TransactionFile) (table.Storage, error) {
	tmpStorage := *storage.Copy()
	events := make([]*events.IEvent, len(trans.Executors))

	for i, exec := range trans.Executors {
		var err error
		// fmt.Println(exec)
		tmpStorage, err = exec.DoExecute(&tmpStorage)

		if err != nil {
			return table.Storage{}, err
		}

		events[i] = exec.ToEvent()
	}

	if transactionLog != nil {
		transactionLog.addTransaction(events, "")
	}

	return tmpStorage, nil
}
