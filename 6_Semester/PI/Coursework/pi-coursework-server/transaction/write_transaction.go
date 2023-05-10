package transaction

import (
	"fmt"
	"pi-coursework-server/events"
	"pi-coursework-server/executor"
	"pi-coursework-server/table"
)

type WriteTransaction struct {
	Executors []executor.IExecutor
}

func NewWriteTransaction() *WriteTransaction {
	return &WriteTransaction{}
}

func (trans *WriteTransaction) Eval(storage table.Storage, transactionLog *TransactionFile) (table.Storage, error) {
	// events := make([]events.IEvent, len(trans.Executors))

	// for i, exec := range trans.Executors {
	// 	tmpStorage, event, err := exec.DoExecute(&nextStorage)
	// 	nextStorage = tmpStorage

	// 	if err != nil {
	// 		return table.Storage{}, err
	// 	}

	// 	// events[i] = exec.ToEvent()
	// 	events[i] = event
	// }

	event := events.NewWriteEvent()

	lastWriteIndex := transactionLog.GetLastWriteIndex()
	// if lastWriteIndex == -1 {
	// 	lastWriteIndex = 0
	// }

	for i := lastWriteIndex + 1; i < len(transactionLog.Logs); i++ {
		fmt.Println("i = ", i)
		Ev := transactionLog.Logs[i].Ev
		// Ev.Apply()
		Ev.GetDescription()
	}
	storage.Save()

	if transactionLog != nil {
		transactionLog.AddSingleEvent(event, "", -1, true)
	}

	return storage, nil
}
