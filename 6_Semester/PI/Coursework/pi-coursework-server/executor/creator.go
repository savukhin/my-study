package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type Creator struct {
	TableName string
	Columns   []string
}

func NewCreator(tableName string, columns []string) *Creator {
	creator := &Creator{
		TableName: tableName,
		Columns:   columns,
	}

	// creator.IExecutor = creator
	return creator
}

func (creator *Creator) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	table, err := table.NewTable(creator.TableName, creator.Columns, make([][]string, 0))
	if err != nil {
		return *copied, nil, err
	}

	// err := table.Save()
	err = copied.AddTable(table)

	return *copied, events.NewCreateEvent(creator.TableName, table.Columns), err
}
