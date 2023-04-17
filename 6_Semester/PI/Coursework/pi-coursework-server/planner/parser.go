package planner

import (
	"errors"
	"regexp"
	"strings"

	"github.com/oriser/regroup"
)

var (
	createTableRegexp = regexp.MustCompile(`^(?i)create\s+table\s+(?P<table_name>\w+)\s+\((?P<cols>(?:\w+,\s*)*\w+)\)$`)
	selectRegexp      = regroup.MustCompile(`(?i)select\s+(?:(?P<columns>(?:(?:\w+,\s*)*\w+))|(?:\*))\s+from\s+(?P<table_name>\w+)(?:\s+(?P<has_where>where)\s+(?P<where_condition>(?P<where_column>\w+)\s*(?P<where_sign>(?:==)|(?:!=))\s+\'(?P<where_value>\w)\'))?(?:\s+(?P<has_limit>limit)\s+(?P<limit>\d+))?`)
	splitRegexp       = regexp.MustCompile(`,\s+`)
)

type Plan struct {
	plan []IProcessor
}

func checkCreateTable(query string) (tableName string, columns []string, err error) {
	result := createTableRegexp.FindStringSubmatch(query)
	if len(result) != 3 {
		err = errors.New("mismatch of columns")
		return
	}

	tableName = result[1]
	columns_str := result[2]
	columns = strings.Split(columns_str, ", ")
	err = nil

	return
}

type WhereCondition struct {
	HasWhere string `regroup:"has_where"`
	Sign     string `regroup:"where_sign"`
	Column   string `regroup:"where_column"`
	Value    string `regroup:"where_value"`
}

type LimitCondition struct {
	Limit    int32  `regroup:"limit"`
	HasLimit string `regroup:"has_limit"`
}

type SelectorGroup struct {
	TableName string `regroup:"table_name"`
	Columns   string `regroup:"columns"`
	Where     WhereCondition
	Limit     LimitCondition
}

func checkSelector(query string) (tableName string, columns []string, whereCondition WhereCondition, limit LimitCondition, err error) {
	elem := &SelectorGroup{}
	err = selectRegexp.MatchToTarget(query, elem)
	if err != nil {
		return
	}

	tableName = elem.TableName
	columns = splitRegexp.Split(elem.Columns, -1)
	whereCondition = elem.Where
	limit = elem.Limit
	return
}

func Parse(query string) (*Plan, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	plan := &Plan{}

	tableName, columns, err := checkCreateTable(query)
	if err == nil {
		plan.plan = append(plan.plan, NewTableCreator(tableName, columns))
		return plan, nil
	}

	// tableName, columns, whereCondition, limit, err := checkSelector(query)
	// if err == nil {
	// 	plan.plan = append(plan.plan, NewTableCreator(tableName, columns))
	// 	return plan, nil
	// }

	return nil, errors.New("no matching pattern")
}
