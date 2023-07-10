package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type MassCreator struct {
	TableName string
	Columns   []string
	Values    [][]string
}

func NewMassCreator(tableName string, columns []string, values [][]string) *MassCreator {
	creator := &MassCreator{
		TableName: tableName,
		Columns:   columns,
		Values:    values,
	}

	// creator.IExecutor = creator
	return creator
}

func (creator *MassCreator) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	table, err := table.NewTable(creator.TableName, creator.Columns, make([][]string, 0))
	if err != nil {
		return *copied, nil, err
	}

	for _, row := range creator.Values {
		table.AddRow(row)
	}

	// err := table.Save()
	err = copied.AddTable(table)

	return *copied, events.NewCreateEvent(creator.TableName, table.Columns), err
}
