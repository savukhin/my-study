package planner

import (
	"strings"

	"github.com/oriser/regroup"
)

var (
	selectRegexp = regroup.MustCompile(`^(?i)select\s+(?:(?P<columns>(?:(?:\w+,\s*)*\w+))|(?:\*))\s+from\s+(?P<table_name>\w+)(?:\s+(?P<has_where>where)\s+(?P<where_condition>(?P<where_column>\w+)\s*(?P<where_sign>(?:==)|(?:!=))\s+\'(?P<where_value>\w)\'))?(?:\s+(?P<has_limit>limit)\s+(?P<limit>\d+))?$`)
)

type WhereCondition struct {
	WhereStr string `regroup:"has_where"`
	Sign     string `regroup:"where_sign"`
	Column   string `regroup:"where_column"`
	Value    string `regroup:"where_value"`
	HasWhere bool
}

type LimitCondition struct {
	Limit    int32  `regroup:"limit"`
	LimitStr string `regroup:"has_limit"`
	HasLimit bool
}

type SelectorGroup struct {
	TableName string `regroup:"table_name"`
	Columns   string `regroup:"columns"`
	Where     WhereCondition
	Limit     LimitCondition
}

func checkSelector(query string) (tableName string, columns []string, whereCondition WhereCondition, limit LimitCondition, err error) {
	elem := &SelectorGroup{}
	err = selectRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	columns = splitRegexp.Split(elem.Columns, -1)

	whereCondition = elem.Where
	whereCondition.HasWhere = (whereCondition.WhereStr != "")

	limit = elem.Limit
	limit.HasLimit = (limit.LimitStr != "")

	return
}
