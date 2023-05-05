package executor

import (
	"errors"
	"fmt"
	"pi-coursework-server/table"
	"pi-coursework-server/utils"
)

type Selector struct {
	TableName string
	Columns   []string
	Where     *utils.WhereConditionCheck
	Limit     *utils.LimitCondition
}

func NewSelector(tableName string, columns []string, where *utils.WhereConditionCheck, limit *utils.LimitCondition) *Selector {
	return &Selector{
		TableName: tableName,
		Columns:   columns,
		Where:     where,
		Limit:     limit,
	}
}

func (selector *Selector) DoExecute(storage *table.Storage) (table.Table, error) {
	tab, err := storage.GetTableCopy(selector.TableName)
	if err != nil {
		return table.Table{}, err
	}

	where := selector.Where
	if where != nil && where.HasWhere {
		result := table.MustNewTable(tab.TableName, tab.Columns, make([][]string, 0))

		if where.Sign != "==" && where.Sign != "!=" {
			return table.Table{}, errors.New("not valid where expression (must be == or != only)")
		}

		column, err := tab.GetColumn(where.Column)
		if err != nil {
			return table.Table{}, err
		}

		val := where.ValueStr

		for i, elem := range column {
			if where.Sign == "==" && elem == val {
				result.AddRow(tab.MustGetRow(i))
			} else if where.Sign == "!=" && elem != val {
				result.AddRow(tab.MustGetRow(i))
			}
		}

		tab = *result
	}

	limit := selector.Limit
	if limit != nil && limit.HasLimit {
		result, err := tab.GetSlice(0, int(limit.Limit))
		if err != nil {
			return table.Table{}, err
		}

		tab = result
	}

	var selectedToDrop []string

	fmt.Println("selector columns", selector.Columns)

	if len(selector.Columns) == 1 && selector.Columns[0] == "*" {
		selectedToDrop = []string{}
	} else {
		selectedToDrop = utils.DifferenceArrays(tab.Columns, selector.Columns)
	}

	fmt.Println("dropping from table", tab, "columns", selectedToDrop)

	return tab.DropColumns(selectedToDrop)
}
