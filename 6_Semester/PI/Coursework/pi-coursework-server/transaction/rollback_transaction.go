package transaction

import (
	"errors"
	"fmt"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type RollbackTransaction struct {
}

func NewRollbackTransaction() *RollbackTransaction {
	return &RollbackTransaction{}
}

func (trans *RollbackTransaction) Eval(storage table.Storage, transactionLog *TransactionFile, transactionName string, transactionTimeUnix int, save bool) (table.Storage, error) {
	nextStorage := *storage.Copy()

	complexTransName, err := transactionLog.GetLastActiveComplexTransactionName()
	fmt.Println("get last active complex transaction err", err, "stack", transactionLog.ActiveComplexTransactions)
	if err != nil {
		return storage, err
	}

	complexTransRollbacked, err := transactionLog.GetRollbackedComplexTransactionByName(complexTransName)
	if err != nil {
		return storage, err
	}

	nextStorage, err = complexTransRollbacked.Eval(nextStorage, nil, transactionName, transactionTimeUnix, true)
	if err != nil {
		return storage, err
	}

	event := events.NewRollbackEvent(complexTransName)
	fmt.Println("Adding rollback event to transaction file")

	transactionLog.AddSingleEvent(event, transactionName, transactionTimeUnix, save)

	return nextStorage, nil
}

func RollbackFromEvent(event events.IEvent) (*RollbackTransaction, error) {
	_, ok := event.(*events.RollbackEvent)
	if !ok {
		return nil, errors.New("not a rollback event")
	}
	return NewRollbackTransaction(), nil
}
