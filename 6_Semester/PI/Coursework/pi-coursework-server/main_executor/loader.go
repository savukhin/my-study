package mainexecutor

import (
	"errors"
	"fmt"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
	"pi-coursework-server/transaction"
)

var (
	transactionFile *transaction.TransactionFile
	cachedStorage   *table.Storage
)

func TryRollback(event events.IEvent, storage *table.Storage, logs *transaction.TransactionFile) (table.Storage, error) {
	_, ok := event.(*events.RollbackEvent)
	if !ok {
		return *storage, errors.New("not a rollback event")
	}

	return transaction.NewRollbackTransaction().Eval(*storage, logs, "", -1, false)
}

func LoadInitialState() error {
	// fmt.Println("Loading storage")
	// currentStorage, err := table.LoadStorage()
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Loading transaction file")
	logs, currentStorage, err := transaction.LoadTransactionFile()
	if err != nil {
		return err
	}

	// --------------------------------------------------------------- //

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

	// --------------------------------------------------------------- //

	// ind := logs.GetLastWriteIndex()
	// if ind == -1 {
	// 	ind = 0
	// }
	// fmt.Println("Loading", len(logs.Logs)-ind, "events")

	// for _, log := range logs.Logs[ind:] {
	// 	// log.
	// 	fmt.Println("Loading event", log.Ev)

	// 	st, err := TryRollback(log.Ev, currentStorage, logs)
	// 	fmt.Println("rollback err", err)
	// 	if err == nil {
	// 		currentStorage = &st
	// 		continue
	// 	}
	// 	// _, ok := log

	// 	exec, err := executor.FromEvent(log.Ev)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println("Executing", log.Ev, "exec =", exec)

	// 	st, _, err = exec.DoExecute(currentStorage)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	currentStorage = &st
	// }

	cachedStorage = currentStorage
	transactionFile = logs

	fmt.Println("Cached storage", cachedStorage)
	musicians, _ := cachedStorage.GetTable("musicians")
	fmt.Println("musicians", musicians)

	return nil
}
