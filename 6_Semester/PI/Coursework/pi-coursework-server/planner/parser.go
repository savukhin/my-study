package planner

import (
	"errors"
	"pi-coursework-server/processors"
	"regexp"
	"strings"
)

var (
	splitRegexp = regexp.MustCompile(`,\s+`)
)

type Plan struct {
	plan []processors.IProcessor
}

func ParseOneString(query string) (*Plan, error) {
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

	tableName, err = checkDropTable(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewDropper(tableName))

		return plan, nil
	}

	tableName, where, err := checkDeleteRows(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))
		plan.plan = append(plan.plan, processors.NewAggregator(where.Column, where.Sign, where.Value))
		plan.plan = append(plan.plan, processors.NewDeleter())

		return plan, nil
	}

	transaction, err := checkBeginTransaction(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewBeginTransaction(transaction))

		return plan, nil
	}

	transaction, err = checkCommitTransaction(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewCommitTransaction(transaction))

		return plan, nil
	}

	err = checkCommit(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewCommit())

		return plan, nil
	}

	err = checkRollback(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewRollback())

		return plan, nil
	}

	return nil, errors.New("no matching pattern")
}
