package executor

import "pi-coursework-server/table"

type Deleter struct {
	IExecutor

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

func (deleter *Deleter) DoExecute(storage *table.Storage) (table.Storage, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(deleter.TableName)
	if err != nil {
		return *copied, err
	}

	columnInd, err := tab.GetColumnIndex(deleter.Column)
	if err != nil {
		return *copied, nil
	}

	y := 0
	for y < tab.Shape.Y {
		row := tab.Values[y]
		if (row[columnInd] == deleter.CompareValues && deleter.Sign == EqualWhereSign) || (row[columnInd] != deleter.CompareValues && deleter.Sign == NotEqualWhereSign) {
			err := tab.DeleteRow(y)
			if err != nil {
				return *copied, err
			}
		} else {
			y++
		}
	}

	return *copied, nil
}
