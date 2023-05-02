package executor

import (
	"errors"
	"pi-coursework-server/table"
)

type Inserter struct {
	IExecutor

	TableName string
	Values    map[string]string
}

func MustNewInserter(tableName string, values, columns []string) *Inserter {
	val, err := NewInserter(tableName, values, columns)
	if err != nil {
		panic(err)
	}
	return val
}

func NewInserter(tableName string, values, columns []string) (*Inserter, error) {
	if len(values) != len(columns) {
		return nil, errors.New("values len doesn't match columns len")
	}

	resultValues := make(map[string]string)
	for i := range values {
		resultValues[columns[i]] = values[i]
	}

	return &Inserter{
		TableName: tableName,
		Values:    resultValues,
	}, nil
}

func NewInserterFromMap(tableName string, Values map[string]string) *Inserter {
	return &Inserter{
		TableName: tableName,
		Values:    Values,
	}
}

func (inserter *Inserter) DoExecute(storage *table.Storage) (table.Storage, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(inserter.TableName)
	if err != nil {
		return *copied, nil
	}

	err = tab.AddRowMap(inserter.Values)
	return *copied, err
}
