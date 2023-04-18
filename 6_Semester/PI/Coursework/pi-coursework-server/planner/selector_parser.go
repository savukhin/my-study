package planner

import (
	"strconv"
	"strings"

	"github.com/oriser/regroup"
)

var (
	selectRegexp = regroup.MustCompile(`(?i)^select\s+(?:(?P<columns>(?:(?:\s*\w+\s*,\s*)*\s*\w+))|(?:\*))\s+from\s+(?P<table_name>\w+)(?:\s+(?P<has_where>where)\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?P<where_value>(?:\'(?P<where_value_str>\w+)\')|(?P<where_value_int>\d+)))?(?:\s+(?P<has_limit>limit)\s+(?P<limit>\d+))?$`)
)

type WhereCondition struct {
	Sign     string `regroup:"where_sign"`
	Column   string `regroup:"where_column"`
	Value    string `regroup:"where_value"`
	ValueStr string `regroup:"where_value_str"`
	ValueInt int32  `regroup:"where_value_int"`
}

func (where *WhereCondition) ExtractValue() string {
	if where.Value[0] == '\'' {
		return where.ValueStr
	} else {
		return strconv.Itoa(int(where.ValueInt))
	}
}

type WhereConditionCheck struct {
	WhereStr string `regroup:"has_where"`
	WhereCondition
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
	Where     WhereConditionCheck
	Limit     LimitCondition
}

func checkSelector(query string) (tableName string, columns []string, whereCondition WhereConditionCheck, limit LimitCondition, err error) {
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
