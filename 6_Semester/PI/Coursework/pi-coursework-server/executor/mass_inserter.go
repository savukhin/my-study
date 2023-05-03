package executor

import (
	"errors"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type MassInserter struct {
	TableName       string
	AfterIndexes    []int
	InsertingValues map[int][]string
}

func NewMassInserter(tableName string, afterIndexes []int, values map[int][]string) (*MassInserter, error) {
	return &MassInserter{
		TableName:       tableName,
		AfterIndexes:    afterIndexes,
		InsertingValues: values,
	}, nil
}

func (massInserter *MassInserter) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(massInserter.TableName)
	if err != nil {
		return *copied, nil, err
	}

	// err = tab.AddRowMap(inserter.Values)
	for _, ind := range massInserter.AfterIndexes {
		values, ok := massInserter.InsertingValues[ind]
		if !ok {
			return *copied, nil, errors.New("data corrupted")
		}

		err = tab.AddRow(values)
	}
	return *copied, nil, err
}
