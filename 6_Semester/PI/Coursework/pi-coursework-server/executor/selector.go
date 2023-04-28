package executor

import (
	"errors"
	"fmt"
	"pi-coursework-server/planner"
	"pi-coursework-server/table"
)

type Selector struct {
	IExecutor

	TableName string
	Where     *planner.WhereConditionCheck
	Limit     *planner.LimitCondition
}

func NewSelector(tableName string, where *planner.WhereConditionCheck, limit *planner.LimitCondition) *Selector {
	return &Selector{
		TableName: tableName,
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
		fmt.Println("LIMIT ERR", err, result.Shape)
		if err != nil {
			return table.Table{}, err
		}

		tab = result
	}

	return tab, nil
}
