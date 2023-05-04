package mainexecutor

import (
	"errors"
	"pi-coursework-server/executor"
	"pi-coursework-server/planner"
	"pi-coursework-server/table"
	"pi-coursework-server/transaction"
	"strings"
)

var (
	transactionFile *transaction.TransactionFile
	cachedStorage   *table.Storage
)

func checkAndTryRollback(query string) error {
	err := planner.CheckRollback(query)
	if err == nil {
		return err
	}

	storage, err := transaction.NewRollbackTransaction().Eval(*cachedStorage, transactionFile)
	if err != nil {
		return err
	}

	cachedStorage = &storage
	return nil
}

func checkAndTryWrite(query string) error {
	storage, err := transaction.NewRollbackTransaction().Eval(*cachedStorage, transactionFile)
	if err != nil {
		return err
	}

	cachedStorage = &storage
	return nil
}

func checkAndTrySelect(query string) (string, error) {
	tableName, columns, whereCondition, limiter, err := planner.CheckSelector(query)
	if err != nil {
		return "", err
	}

	exec := executor.NewSelector(tableName, columns, &whereCondition, &limiter)
	// exec.Columns
	table, err := exec.DoExecute(cachedStorage)
	if err != nil {
		return "", err
	}

	return table.ToString(), nil
}

func checkAndTrySingleLineCommand(query string) (string, error) {
	if err := checkAndTryRollback(query); err == nil {
		return "ok", nil
	}

	if err := checkAndTryWrite(query); err == nil {
		return "ok", nil
	}

	if str, err := checkAndTrySelect(query); err == nil {
		return str, nil
	}

	return "", errors.New("single-line command not recognized")
}

func ExecuteWholeQuery(query string) (string, error) {
	query = strings.Replace(query, "\n", " ", -1)
	commands := strings.Split(query, ";")

	if len(commands) == 1 {
		command := strings.Trim(commands[0], " \t\n")
		return checkAndTrySingleLineCommand(command)
	}

	// transactionBlocks := make([][]transaction.ITransaction, 0)
	currentBlock := make([]executor.IExecutor, 0)
	wasBegin := false

	response := make([]string, 0)

	for _, command := range commands {
		command = strings.Trim(command, " \t\n")

		if command == "" {
			continue
		}

		if len(currentBlock) == 0 && !wasBegin {
			str, err := checkAndTrySingleLineCommand(query)
			if err == nil {
				response = append(response, str)
				continue
			}

			err = planner.CheckBegin(query)
			if err == nil {
				wasBegin = true
				continue
			}

			return "", errors.New("no begin in transaction")
		}

		if err := planner.CheckBegin(query); err == nil {
			return "", errors.New("more than one begin in transaction")
		}

		if err := planner.CheckCommit(query); err == nil {
			complex := transaction.NewComplexTransaction(currentBlock)
			storage2, err := complex.Eval(*cachedStorage, transactionFile)
			if err != nil {
				return "", err
			}
			cachedStorage = &storage2

			currentBlock = make([]executor.IExecutor, 0)
		}

		exec, err := planner.ParseOneQueryForExecutor(command)
		if err != nil {
			return "", err
		}

		currentBlock = append(currentBlock, exec)
	}

	result := ""
	for _, output := range response {
		result = result + "\n\n" + output
	}

	return result, nil
}
