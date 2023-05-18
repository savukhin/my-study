package planner

import (
	"errors"
	"pi-coursework-server/executor"
	"strings"
)

func ParseOneQueryForExecutor(query string) (executor.IExecutor, error) {
	query = strings.Trim(query, " \t\n")
	// query_lowed := strings.ToLower(query)

	if query == "" {
		return nil, nil
	}

	username, password, err := CheckAddUser(query)
	if err == nil {
		exec := executor.NewInserterFromMap(PIDBUsersTable, map[string]string{
			"username": username,
			"password": password,
		})

		return exec, nil
	}

	tableName, columns, err := CheckCreateTable(query)
	if err == nil {
		exec := executor.NewCreator(tableName, columns)

		return exec, nil
	}

	tableName, values, err := CheckInsert(query)
	if err == nil {
		exec := executor.NewInserterFromMap(tableName, values)

		return exec, nil
	}

	tableName, setColumnName, setValue, where, err := CheckUpdate(query)
	if err == nil {
		exec, err := executor.NewUpdater(tableName,
			where.Column, executor.WhereSign(where.Sign), where.Value,
			map[string]string{
				setColumnName: setValue,
			})

		return exec, err
	}

	tableName, err = CheckDropTable(query)
	if err == nil {
		exec := executor.NewDropper(tableName)

		return exec, nil
	}

	tableName, where, err = CheckDeleteRows(query)
	if err == nil {
		// plan.Plan = append(plan.Plan, processors.NewTableGetter(tableName))
		// plan.Plan = append(plan.Plan, processors.NewAggregator(where.Column, where.Sign, where.ExtractValue()))
		// plan.Plan = append(plan.Plan, processors.NewDeleter())
		exec := executor.NewDeleter(tableName, where.Column, executor.WhereSign(where.Sign), where.Value)

		return exec, nil
	}

	// 	transaction, err := CheckBeginTransaction(query)
	// 	if err == nil {
	// 		plan.Plan = append(plan.Plan, processors.NewBeginTransaction(transaction))

	// 		return plan, nil
	// 	}

	// 	transaction, err = CheckCommitTransaction(query)
	// 	if err == nil {
	// 		plan.Plan = append(plan.Plan, processors.NewCommitTransaction(transaction))

	// 		return plan, nil
	// 	}

	// 	err = CheckCommit(query)
	// 	if err == nil {
	// 		plan.Plan = append(plan.Plan, processors.NewCommit())

	// 		return plan, nil
	// 	}

	// 	err = CheckRollback(query)
	// 	if err == nil {
	// 		plan.Plan = append(plan.Plan, processors.NewRollback())

	// 		return plan, nil
	// 	}

	// 	return nil, errors.New("no matching pattern")
	// }

	// func ParseFullQuery(query string) (*Plan, error) {
	// 	query = strings.Replace(query, "\n", " ", -1)
	// 	commands := strings.Split(query, ";")

	// 	plan := NewPlan()

	// 	for _, command := range commands {
	// 		subPlan, err := ParseOneString(command)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 		plan.Plan = append(plan.Plan, subPlan.Plan...)
	// 	}

	// return plan, nil

	return nil, errors.New("command not recognized")
}
