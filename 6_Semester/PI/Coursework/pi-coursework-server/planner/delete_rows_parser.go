package planner

import (
	"fmt"
	"pi-coursework-server/utils"
	"strings"

	"github.com/oriser/regroup"
)

var (
	deleteRowsRegexp = regroup.MustCompile(`(?i)delete\s+from\s+(?P<table_name>\w+)\s+where\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?:\'(?P<where_value>\w+)\')`)
)

type DeleteRowsGroup struct {
	TableName string `regroup:"table_name"`
	utils.WhereCondition
}

func CheckDeleteRows(query string) (tableName string, where utils.WhereCondition, err error) {
	elem := &DeleteRowsGroup{}
	err = deleteRowsRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	fmt.Println("Checking delete rows ", query, "err ", err)
	if err != nil {
		return
	}

	tableName = elem.TableName
	where = elem.WhereCondition
	err = nil

	return
}
