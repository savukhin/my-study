package planner

import (
	"pi-coursework-server/utils"
	"strings"

	"github.com/oriser/regroup"
)

var (
	deleteRowsRegexp = regroup.MustCompile(`(?i)delete\s+from\s+(?P<table_name>\w+)\s+where\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?P<where_value>(?:\'(?P<where_value_str>\w+)\')|(?P<where_value_int>\d+))`)
)

type DeleteRowsGroup struct {
	TableName string `regroup:"table_name"`
	utils.WhereCondition
}

func CheckDeleteRows(query string) (tableName string, where utils.WhereCondition, err error) {
	elem := &DeleteRowsGroup{}
	err = deleteRowsRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	where = elem.WhereCondition
	err = nil

	return
}
