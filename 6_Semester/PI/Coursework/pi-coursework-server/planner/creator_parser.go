package planner

import (
	"strings"

	"github.com/oriser/regroup"
)

var (
	createTableRegexp = regroup.MustCompile(`^(?i)create\s+table\s+(?P<table_name>\w+)\s+\((?P<columns>(?:\w+,\s*)*\w+)\)$`)
)

type CreateTableGroup struct {
	TableName string `regroup:"table_name"`
	Columns   string `regroup:"columns"`
}

func checkCreateTable(query string) (tableName string, columns []string, err error) {
	elem := &CreateTableGroup{}
	err = createTableRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	columns = strings.Split(elem.Columns, ", ")
	err = nil

	return
}
