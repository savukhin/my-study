package planner

import (
	"pi-coursework-server/utils"
	"strings"

	"github.com/oriser/regroup"
)

var (
	updateRegexp = regroup.MustCompile(`(?i)^update\s+(?P<table_name>\w+)\s+set\s+(?P<set_column_name>\w+)\s*=\s*\'(?P<set_value>\w+)\'\s+where\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?P<where_value>(?:\'(?P<where_value_str>\w+)\')|(?P<where_value_int>\d+))$`)
)

type UpdateGroup struct {
	TableName     string `regroup:"table_name"`
	SetColumnName string `regroup:"set_column_name"`
	SetValue      string `regroup:"set_value"`
	utils.WhereCondition
}

func CheckUpdate(query string) (tableName, setColumnName, setValue string, where utils.WhereCondition, err error) {
	elem := &UpdateGroup{}
	err = updateRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	setColumnName = elem.SetColumnName
	setValue = elem.SetValue
	where = elem.WhereCondition
	err = nil

	return
}
