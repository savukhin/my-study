package mainexecutor

import (
	"errors"
	"fmt"
	"pi-coursework-server/executor"
	"pi-coursework-server/planner"
	"pi-coursework-server/transaction"
	"strings"
)

func checkAndTryRollback(query string) (bool, error) {
	err := planner.CheckRollback(query)
	if err != nil {
		return false, err
	}

	fmt.Println("Catched rollback in query", query)

	storage, err := transaction.NewRollbackTransaction().Eval(*cachedStorage, transactionFile, "", -1, true)
	if err != nil {
		return true, err
	}

	cachedStorage = &storage
	return true, nil
}

func checkAndTryWrite(query string) (bool, error) {
	err := planner.CheckWrite(query)
	if err != nil {
		return false, err
	}
	fmt.Println("Catched write in query", query)

	storage, err := transaction.NewWriteTransaction().Eval(*cachedStorage, transactionFile)
	if err != nil {
		return true, err
	}

	cachedStorage = &storage
	return true, nil
}

func checkAndTrySelect(query string) (string, bool, error) {
	tableName, columns, whereCondition, limiter, err := planner.CheckSelector(query)
	fmt.Println("check select", err)
	if err != nil {
		return "", false, err
	}

	fmt.Println("Catched selector in query", query)

	exec := executor.NewSelector(tableName, columns, &whereCondition, &limiter)
	// exec.Columns
	table, err := exec.DoExecute(cachedStorage)
	fmt.Println("err", err)
	if err != nil {
		return "", true, err
	}

	return table.ToString(), false, nil
}

func checkAndTrySingleLineCommand(query string) (string, bool, error) {
	checked, err := checkAndTryRollback(query)
	if err == nil {
		return "ok", checked, nil
	} else if checked {
		return "err", checked, err
	}

	checked, err = checkAndTryWrite(query)
	if err == nil {
		return "ok", checked, nil
	} else if checked {
		return "err", checked, err
	}

	str, checked, err := checkAndTrySelect(query)
	if err == nil {
		return str, checked, nil
	} else if checked {
		return "err", checked, err
	}

	return "", false, errors.New("single-line command not recognized")
}

func ExecuteWholeQuery(query string) ([]string, error) {
	query = strings.Replace(query, "\n", " ", -1)
	commands := strings.Split(query, ";")

	fmt.Println("commands", commands)

	// if len(commands) == 1 {
	// 	command := strings.Trim(commands[0], " \t\n")
	// 	return checkAndTrySingleLineCommand(command)
	// }

	// transactionBlocks := make([][]transaction.ITransaction, 0)
	currentBlock := make([]executor.IExecutor, 0)
	wasBegin := false

	responses := make([]string, 0)

	for _, command := range commands {
		fmt.Println(command)

		command = strings.Trim(command, " \t\n")

		if command == "" {
			continue
		}

		if len(currentBlock) == 0 && !wasBegin {
			str, checked, err := checkAndTrySingleLineCommand(command)
			if checked && err != nil {
				return responses, err
			}
			if err == nil {
				responses = append(responses, str)
				continue
			}

			err = planner.CheckBegin(command)
			if err == nil {
				wasBegin = true
				continue
			}

			return responses, errors.New("no begin in transaction")
		}

		if err := planner.CheckBegin(command); err == nil {
			return responses, errors.New("more than one begin in transaction")
		}

		if err := planner.CheckCommit(command); err == nil {
			complex := transaction.NewComplexTransaction(currentBlock)
			fmt.Println("Catched complex in query", currentBlock)
			storage2, err := complex.Eval(*cachedStorage, transactionFile, "", -1, true)
			if err != nil {
				return responses, err
			}
			cachedStorage = &storage2

			currentBlock = make([]executor.IExecutor, 0)
			wasBegin = false

			responses = append(responses, "ok")

			continue
		}

		exec, err := planner.ParseOneQueryForExecutor(command)
		if err != nil {
			return responses, err
		}

		currentBlock = append(currentBlock, exec)
	}

	return responses, nil
}
