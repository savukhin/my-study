package transaction

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type RollbackTransaction struct {
}

func NewRollbackTransaction() *RollbackTransaction {
	return &RollbackTransaction{}
}

func (trans *RollbackTransaction) Eval(storage table.Storage, transactionLog *TransactionFile) (table.Storage, error) {
	nextStorage := *storage.Copy()

	complexTransName, err := transactionLog.GetLastActiveComplexTransactionName()
	if err != nil {
		return storage, nil
	}

	complexTransRollbacked, err := transactionLog.GetRollbackedComplexTransactionByName(complexTransName)
	if err != nil {
		return storage, err
	}

	nextStorage, err = complexTransRollbacked.Eval(nextStorage, nil)
	if err != nil {
		return storage, err
	}

	event := events.NewRollbackEvent(complexTransName)

	transactionLog.AddSingleEvent(event, "")

	return nextStorage, nil
}
