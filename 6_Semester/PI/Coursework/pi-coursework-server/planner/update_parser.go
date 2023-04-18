package planner

import (
	"strings"

	"github.com/oriser/regroup"
)

var (
	updateRegexp = regroup.MustCompile(`^update\s+(?P<table_name>\w+)\s+set\s+(?P<set_column_name>\w+)\s*=\s*\'(?P<set_value>\w+)\'\s+where\s+(?P<where_column>\w+)\s+(?P<where_sign>(?:==)|(?:!=))\s+(?P<where_value>(?:\'(?P<where_value_str>\w+)\')|(?P<where_value_int>\d+))$`)
)

type UpdateGroup struct {
	TransactionName string `regroup:"transaction_name"`
}

func checkUpdate(query string) (transactionName string, err error) {
	elem := &TransactionGroup{}
	err = beginTransactionRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	transactionName = elem.TransactionName
	err = nil

	return
}
