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

	fmt.Println("Loading transaction file")
	logs, err := transaction.LoadTransactionFile()
	if err != nil {
		return err
	}

	// complexTransName, err := transactionFile.GetLastActiveComplexTransactionName()
	// if err != nil {
	// 	return err
	// }

	// complexTransRollbacked, err := transactionFile.GetComplexTransactionByName(complexTransName)
	// if err != nil {
	// 	return err
	// }

	// cachedStorage2, err = complexTransRollbacked.Eval(*cachedStorage2, nil)
	// if err != nil {
	// 	return err
	// }

	// event := events.NewRollbackEvent(complexTransName)

	// transactionLog.AddSingleEvent(event, "")

	ind := logs.GetLastWriteIndex()
	if ind == -1 {
		ind = 0
	}
	fmt.Println("Loading", len(logs.Logs)-ind, "events")
	for _, log := range logs.Logs[ind:] {
		// log.
		exec, err := executor.FromEvent(log.Ev)
		if err != nil {
			return err
		}

		fmt.Println("Executing", log.Ev, "exec =", exec)

		st, _, err := exec.DoExecute(cachedStorage2)
		if err != nil {
			return err
		}
		cachedStorage2 = &st
	}

	cachedStorage = cachedStorage2
	transactionFile = logs

	fmt.Println("Cached storage", cachedStorage)
	musicians, _ := cachedStorage.GetTable("musicians")
	fmt.Println("musicians", musicians)

	return nil
}

func checkAndTryRollback(query string) (bool, error) {
	err := planner.CheckRollback(query)
	if err != nil {
		return false, err
	}

	fmt.Println("Catched rollback in query", query)

	storage, err := transaction.NewRollbackTransaction().Eval(*cachedStorage, transactionFile)
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
	if checked, err := checkAndTryRollback(query); err == nil {
		return "ok", checked, nil
	}

	if checked, err := checkAndTryWrite(query); err == nil {
		return "ok", checked, nil
	}

	if str, checked, err := checkAndTrySelect(query); err == nil {
		return str, checked, nil
	}

	return "", false, errors.New("single-line command not recognized")
}

func ExecuteWholeQuery(query string) (string, error) {
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

	response := make([]string, 0)

	for _, command := range commands {
		fmt.Println(command)

		command = strings.Trim(command, " \t\n")

		if command == "" {
			continue
		}

		if len(currentBlock) == 0 && !wasBegin {
			str, checked, err := checkAndTrySingleLineCommand(command)
			if checked && err != nil {
				return "", err
			}
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
		if result == "" {
			result = output
			continue
		}
		if output == "" {
			continue
		}
		result = result + "\n\n" + output
	}

	return result, nil
}
