package transaction

import (
	"errors"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type ITransaction interface {
	Eval(table.Storage, *TransactionFile, string, int, bool) (table.Storage, error)
}

// type AbstractTransaction struct {
// 	ITransaction
// }

func FromEvents(eventsArr []events.IEvent) (ITransaction, error) {
	complex, err := ComplexFromEvents(eventsArr)
	if err == nil {
		return complex, nil
	}

	if len(eventsArr) != 1 {
		return nil, errors.New("no such transaction for this event")
	}

	rollback, err := RollbackFromEvent(eventsArr[0])
	if err == nil {
		return rollback, nil
	}

	return nil, errors.New("no such transaction for event block")
}
