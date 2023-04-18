package planner

import (
	"errors"
	"pi-coursework-server/processors"
	"regexp"
	"strings"

	"github.com/oriser/regroup"
)

var (
	dropTableRegexp         = regroup.MustCompile(`(?i)drop\s+table\s+(?P<table_name>\w+)`)
	deleteRowsRegexp        = regroup.MustCompile(`(?i)delete\s+from\s+(?P<table_name>\w+)\s+where\s+(?P<where_column>\w+)\s*(?P<where_sign>(?:==)|(?:!=))\s+\'(?P<where_value>\w)\'`)
	beginTransactionRegexp  = regroup.MustCompile(`(?i)begin\s+(?P<transaction_name>\w+)`)
	commitTransactionRegexp = regroup.MustCompile(`(?i)commit\s+(?P<transaction_name>\w+)`)
	commitRegexp            = regroup.MustCompile(`(?i)commit`)
	rollbackRegexp          = regroup.MustCompile(`(?i)rollback`)
	splitRegexp             = regexp.MustCompile(`,\s+`)
)

type Plan struct {
	plan []processors.IProcessor
}

func Parse(query string) (*Plan, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	plan := &Plan{}

	tableName, columns, err := checkCreateTable(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableCreator(tableName, columns))
		return plan, nil
	}

	tableName, columns, whereCondition, limiter, err := checkSelector(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))

		if whereCondition.HasWhere {
			plan.plan = append(plan.plan, processors.NewAggregator(whereCondition.Column, whereCondition.Sign, whereCondition.Value))
		}

		plan.plan = append(plan.plan, processors.NewSelector(columns))

		if limiter.HasLimit {
			plan.plan = append(plan.plan, processors.NewLimiter(limiter.Limit))
		}

		return plan, nil
	}

	return nil, errors.New("no matching pattern")
}
