package planner

import (
	"errors"
	"regexp"
	"strings"
)

var (
	createTableRegexp = regexp.MustCompile(`^(?i)create\s+table\s+(?P<name>\w+)\s+\((?P<cols>(?:\w+,\s*)*\w+)\)$`)
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

func Parse(query string) (*Plan, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	plan := &Plan{}

	tableName, columns, err := checkCreateTable(query)
	if err == nil {
		plan.plan = append(plan.plan, NewTableCreator(tableName, columns))
		return plan, nil
	}

	return nil, errors.New("no matching pattern")
}
