package planner

import (
	"errors"
	"regexp"
	"strings"

	"github.com/oriser/regroup"
)

var (
	beginTransactionRegexp  = regroup.MustCompile(`^(?i)begin\s+(?P<transaction_name>\w+)$`)
	commitTransactionRegexp = regroup.MustCompile(`^(?i)commit\s+(?P<transaction_name>\w+)$`)
	commitRegexp            = regexp.MustCompile(`^(?i)commit$`)
	rollbackRegexp          = regexp.MustCompile(`^(?i)rollback$`)
)

type TransactionGroup struct {
	TransactionName string `regroup:"transaction_name"`
}

func checkBeginTransaction(query string) (transactionName string, err error) {
	elem := &TransactionGroup{}
	err = beginTransactionRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	transactionName = elem.TransactionName
	err = nil

	return
}

func checkCommitNamedTransaction(query string) (transactionName string, err error) {
	elem := &TransactionGroup{}
	err = commitTransactionRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	transactionName = elem.TransactionName
	err = nil

	return
}

func checkCommitTransaction(query string) (err error) {
	matched := commitRegexp.MatchString(strings.TrimSpace(query))
	if !matched {
		err = errors.New("not matched")
		return
	}
	err = nil

	return
}

func checkRollback(query string) (err error) {
	matched := rollbackRegexp.MatchString(strings.TrimSpace(query))
	if !matched {
		err = errors.New("not matched")
		return
	}
	err = nil

	return
}