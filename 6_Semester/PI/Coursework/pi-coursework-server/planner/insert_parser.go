package planner

import (
	"errors"
	"strings"

	"github.com/oriser/regroup"
)

var (
	insertRegexp = regroup.MustCompile(`(?i)^insert\s+into\s+(?P<table_name>\w+)\s*\((?P<columns>(?:(?:\s*\w+\s*,\s*)*\s*\w+\s*))\)\s+values\s*\((?P<values>(?:(?:\s*\w+\s*,\s*)*\s*\w+\s*))\)$`)
)

type InsertGroup struct {
	TableName string `regroup:"table_name"`
	Columns   string `regroup:"columns"`
	Values    string `regroup:"values"`
}

func CheckInsert(query string) (tableName string, values map[string]string, err error) {
	elem := &InsertGroup{}
	err = insertRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	tableName = elem.TableName

	columns := splitRegexp.Split(strings.TrimSpace(elem.Columns), -1)
	elems := splitRegexp.Split(strings.TrimSpace(elem.Values), -1)
	if len(columns) != len(elems) {
		err = errors.New("elems lenght not equal to columns length")
		return
	}

	values = make(map[string]string)
	for i, column := range columns {
		values[column] = elems[i]
	}

	err = nil

	return
}
