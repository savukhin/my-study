package mainexecutor

import (
	"errors"
	"fmt"
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

func LoadInitialState() error {
	fmt.Println("Loading storage")
	cachedStorage2, err := table.LoadStorage()
	if err != nil {
		return err
	}
	cachedStorage = cachedStorage2

	fmt.Println("Loading transaction file")
	log, err := transaction.LoadTransactionFile()
	if err != nil {
		return err
	}
	transactionFile = log

	return nil
}

func checkAndTryRollback(query string) error {
	err := planner.CheckRollback(query)
	if err != nil {
		return err
	}

	fmt.Println("Catched rollback in query", query)

	storage, err := transaction.NewRollbackTransaction().Eval(*cachedStorage, transactionFile)
	if err != nil {
		return err
	}

	cachedStorage = &storage
	return nil
}

func checkAndTryWrite(query string) error {
	err := planner.CheckWrite(query)
	if err != nil {
		return err
	}
	fmt.Println("Catched write in query", query)

	storage, err := transaction.NewWriteTransaction().Eval(*cachedStorage, transactionFile)
	if err != nil {
		return err
	}

	cachedStorage = &storage
	return nil
}

func checkAndTrySelect(query string) (string, error) {
	tableName, columns, whereCondition, limiter, err := planner.CheckSelector(query)
	fmt.Println("check select", err)
	if err != nil {
		return "", err
	}

	fmt.Println("Catched selector in query", query)

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
		fmt.Println(command)

		command = strings.Trim(command, " \t\n")

		if command == "" {
			continue
		}

		if len(currentBlock) == 0 && !wasBegin {
			str, err := checkAndTrySingleLineCommand(command)
			if err == nil {
				response = append(response, str)
				continue
			}

			err = planner.CheckBegin(command)
			if err == nil {
				wasBegin = true
				continue
			}

			return "", errors.New("no begin in transaction")
		}

		if err := planner.CheckBegin(command); err == nil {
			return "", errors.New("more than one begin in transaction")
		}

		if err := planner.CheckCommit(command); err == nil {
			complex := transaction.NewComplexTransaction(currentBlock)
			fmt.Println("Catched complex in query", currentBlock)
			storage2, err := complex.Eval(*cachedStorage, transactionFile)
			if err != nil {
				return "", err
			}
			cachedStorage = &storage2

			currentBlock = make([]executor.IExecutor, 0)
			wasBegin = false

			continue
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
