package planner

import (
	"errors"
	"pi-coursework-server/processors"
	"regexp"
	"strings"
)

const (
	PIDBUsersTable = "pidb_users"
)

var (
	splitRegexp = regexp.MustCompile(`\s*,\s*`)
)

type Plan struct {
	plan []processors.IProcessor
}

func NewPlan() *Plan {
	return &Plan{
		plan: make([]processors.IProcessor, 0),
	}
}

func ParseOneString(query string) (*Plan, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	plan := NewPlan()

	if query == "" {
		return plan, nil
	}
	username, password, err := checkAddUser(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(PIDBUsersTable))
		plan.plan = append(plan.plan, processors.NewInserter(map[string]string{"username": username, "password": password}))

		return plan, nil
	}

	tableName, columns, err := checkCreateTable(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableCreator(tableName, columns))
		return plan, nil
	}

	tableName, values, err := checkInsert(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))
		plan.plan = append(plan.plan, processors.NewInserter(values))

		return plan, nil
	}

	tableName, setColumnName, setValue, where, err := checkUpdate(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))
		plan.plan = append(plan.plan, processors.NewAggregator(where.Column, where.Sign, where.ExtractValue()))
		plan.plan = append(plan.plan, processors.NewUpdater(setColumnName, setValue))

		return plan, nil
	}

	tableName, columns, whereCondition, limiter, err := checkSelector(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))

		if whereCondition.HasWhere {
			plan.plan = append(plan.plan, processors.NewAggregator(whereCondition.Column, whereCondition.Sign, whereCondition.ExtractValue()))
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

	tableName, where, err = checkDeleteRows(query)
	if err == nil {
		plan.plan = append(plan.plan, processors.NewTableGetter(tableName))
		plan.plan = append(plan.plan, processors.NewAggregator(where.Column, where.Sign, where.ExtractValue()))
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

func ParseFullQuery(query string) (*Plan, error) {
	query = strings.Replace(query, "\n", " ", -1)
	commands := strings.Split(query, ";")

	plan := NewPlan()

	for _, command := range commands {
		subPlan, err := ParseOneString(command)
		if err != nil {
			return nil, err
		}

		plan.plan = append(plan.plan, subPlan.plan...)
	}

	return plan, nil
}
