package executor

import (
	"pi-coursework-server/planner"
	"pi-coursework-server/table"
	"pi-coursework-server/transaction"
)

type IExecutor interface {
	DoExecute(*table.Storage, *transaction.TransactionFile) (table.Storage, error)
	DoExecuteCallback(table.Storage) error
}

func ExecuteQuery(storage *table.Storage, query string) (*table.Table, error) {
	plan, err := planner.ParseFullQuery(query)
	if err != nil {
		return nil, err
	}

	// checkpoint := new(table.Storage)
	checkpoint := *storage

	for _, processor := range plan.Plan {
		_, err := processor.DoProcess(checkpoint)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
