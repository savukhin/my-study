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

func ParseOneString(query string) (*Plan, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	plan := NewPlan()

	if query == "" {
		return plan, nil
	}
	username, password, err := CheckAddUser(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableGetter(PIDBUsersTable))
		plan.Plan = append(plan.Plan, processors.NewInserter(map[string]string{"username": username, "password": password}))

		return plan, nil
	}

	tableName, columns, err := CheckCreateTable(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableCreator(tableName, columns))
		return plan, nil
	}

	tableName, values, err := CheckInsert(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableGetter(tableName))
		plan.Plan = append(plan.Plan, processors.NewInserter(values))

		return plan, nil
	}

	tableName, setColumnName, setValue, where, err := CheckUpdate(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableGetter(tableName))
		plan.Plan = append(plan.Plan, processors.NewAggregator(where.Column, where.Sign, where.Value))
		plan.Plan = append(plan.Plan, processors.NewUpdater(setColumnName, setValue))

		return plan, nil
	}

	tableName, columns, whereCondition, limiter, err := CheckSelector(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableGetter(tableName))

		if whereCondition.HasWhere {
			plan.Plan = append(plan.Plan, processors.NewAggregator(whereCondition.Column, whereCondition.Sign, whereCondition.Value))
		}

		plan.Plan = append(plan.Plan, processors.NewSelector(columns))

		if limiter.HasLimit {
			plan.Plan = append(plan.Plan, processors.NewLimiter(limiter.Limit))
		}

		return plan, nil
	}

	tableName, err = CheckDropTable(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewDropper(tableName))

		return plan, nil
	}

	tableName, where, err = CheckDeleteRows(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewTableGetter(tableName))
		plan.Plan = append(plan.Plan, processors.NewAggregator(where.Column, where.Sign, where.Value))
		plan.Plan = append(plan.Plan, processors.NewDeleter())

		return plan, nil
	}

	transaction, err := CheckBeginTransaction(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewBeginTransaction(transaction))

		return plan, nil
	}

	transaction, err = CheckCommitTransaction(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewCommitTransaction(transaction))

		return plan, nil
	}

	err = CheckCommit(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewCommit())

		return plan, nil
	}

	err = CheckRollback(query)
	if err == nil {
		plan.Plan = append(plan.Plan, processors.NewRollback())

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

		plan.Plan = append(plan.Plan, subPlan.Plan...)
	}

	return plan, nil
}
