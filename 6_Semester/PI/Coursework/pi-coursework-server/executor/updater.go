package executor

import (
	"errors"
	"pi-coursework-server/table"
)

type WhereSign string

const (
	EqualWhereSign    WhereSign = "=="
	NotEqualWhereSign WhereSign = "!="
)

type Updater struct {
	IExecutor

	TableName     string
	Column        string
	Sign          WhereSign
	CompareValues string
	NewValues     map[string]string
}

func NewUpdater(tableName, column string, sign WhereSign, compareValue string, newValues map[string]string) (*Updater, error) {
	if sign != "==" && sign != "!=" {
		return nil, errors.New("not valid sign")
	}

	return &Updater{
		TableName:     tableName,
		Column:        column,
		Sign:          sign,
		CompareValues: compareValue,
		NewValues:     newValues,
	}, nil
}

func (updater *Updater) DoExecute(storage *table.Storage) (table.Storage, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(updater.TableName)
	if err != nil {
		return *copied, err
	}

	columnInd, err := tab.GetColumnIndex(updater.Column)
	if err != nil {
		return *copied, nil
	}

	for x, row := range tab.Values {
		if row[columnInd] == updater.CompareValues && updater.Sign == EqualWhereSign {
			err := tab.UpdateRow(x, updater.NewValues)
			if err != nil {
				return *copied, err
			}
		}
	}

	return *copied, nil
}
