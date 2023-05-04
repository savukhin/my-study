package planner

import (
	"pi-coursework-server/utils"
	"strings"

	"github.com/oriser/regroup"
)

var (
	selectRegexp = regroup.MustCompile(`(?i)^select\s+(?:(?P<columns>(?:(?:\s*\w+\s*,\s*)*\s*\w+))|(?:\*))\s+from\s+(?P<table_name>\w+)(?:\s+(?P<has_where>where)\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?P<where_value>(?:\'(?P<where_value_str>\w+)\')|(?P<where_value_int>\d+)))?(?:\s+(?P<has_limit>limit)\s+(?P<limit>\d+))?$`)
)

type SelectorGroup struct {
	TableName string `regroup:"table_name"`
	Columns   string `regroup:"columns"`
	Where     utils.WhereConditionCheck
	Limit     utils.LimitCondition
}

func CheckSelector(query string) (tableName string, columns []string, whereCondition utils.WhereConditionCheck, limit utils.LimitCondition, err error) {
	elem := &SelectorGroup{}
	err = selectRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	columns = splitRegexp.Split(strings.TrimSpace(elem.Columns), -1)

	whereCondition = elem.Where
	whereCondition.HasWhere = (whereCondition.WhereStr != "")

	limit = elem.Limit
	limit.HasLimit = (limit.LimitStr != "")

	return
}
