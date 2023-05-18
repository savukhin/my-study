package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type Deleter struct {
	TableName     string
	Column        string
	Sign          WhereSign
	CompareValues string
}

func NewDeleter(tableName string, column string, sign WhereSign, compareValue string) *Deleter {
	return &Deleter{
		TableName:     tableName,
		Column:        column,
		Sign:          sign,
		CompareValues: compareValue,
	}
}

func (deleter *Deleter) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(deleter.TableName)
	if err != nil {
		return *copied, nil, err
	}

	columnInd, err := tab.GetColumnIndex(deleter.Column)
	if err != nil {
		return *copied, nil, err
	}

	yDeleted := make([]int, 0)
	deletedValues := make(map[int][]string)

	y := 0
	for y < tab.Shape.Y {
		row := tab.Values[y]
		if (row[columnInd] == deleter.CompareValues && deleter.Sign == EqualWhereSign) || (row[columnInd] != deleter.CompareValues && deleter.Sign == NotEqualWhereSign) {
			yAbsolute := y + len(yDeleted)
			yDeleted = append(yDeleted, yAbsolute)

			deletedValues[yAbsolute] = row

			err := tab.DeleteRow(y)

			if err != nil {
				return *copied, nil, err
			}
		} else {
			y++
		}
	}

	event := events.NewDeleteEvent(tab.TableName, yDeleted, deletedValues)

	return *copied, event, nil
}
