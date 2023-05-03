package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/planner"
	"pi-coursework-server/table"
)

type IExecutor interface {
	DoExecute(*table.Storage) (table.Storage, error)
	ToEvent() *events.IEvent
	// DoExecuteCallback(table.Storage) error
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
