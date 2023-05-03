package executor

import (
	"errors"
	"pi-coursework-server/events"
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

func MustNewUpdater(tableName, column string, sign WhereSign, compareValue string, newValues map[string]string) *Updater {
	val, err := NewUpdater(tableName, column, sign, compareValue, newValues)
	if err != nil {
		panic(err)
	}
	return val
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

func (updater *Updater) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(updater.TableName)
	if err != nil {
		return *copied, nil, err
	}

	columnInd, err := tab.GetColumnIndex(updater.Column)
	if err != nil {
		return *copied, nil, err
	}

	yUpdate := make([]int, 0)
	oldValues := make(map[int][]string)

	for y, row := range tab.Values {
		if (row[columnInd] == updater.CompareValues && updater.Sign == EqualWhereSign) || (row[columnInd] != updater.CompareValues && updater.Sign == NotEqualWhereSign) {
			yUpdate = append(yUpdate, y)
			oldValues[y] = make([]string, len(row))
			copy(oldValues[y], row)

			err := tab.UpdateRow(y, updater.NewValues)
			if err != nil {
				return *copied, nil, err
			}
		}
	}

	return *copied, events.NewUpdateEvent(tab.TableName, yUpdate, updater.NewValues, oldValues), nil
}
