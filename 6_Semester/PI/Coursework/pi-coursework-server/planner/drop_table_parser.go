package planner

import (
	"strings"

	"github.com/oriser/regroup"
)

var (
	dropTableRegexp = regroup.MustCompile(`^(?i)drop\s+table\s+(?P<table_name>\w+)$`)
)

type DropTableGroup struct {
	TableName string `regroup:"table_name"`
}

func checkDropTable(query string) (tableName string, err error) {
	elem := &DropTableGroup{}
	err = dropTableRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	err = nil

	return
}
