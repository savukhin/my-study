package transaction

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
	"path"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
	"pi-coursework-server/utils"
	"strconv"
)

var (
	TRANSACATION_FOLDER_PATH = utils.GetEnvDefault("TRANSACATION_FOLDER_PATH", path.Join(".", "static", "transactions"))
	// TRANSACATION_FILE_PATH = utils.GetEnvDefault("TRANSACATION_FILE_PATH", path.Join(".", "static", "transactions", "transactions.csv"))
	columns = []string{
		"time",
		"transaction_name",
		"event_type",
		// "event_table",
		"event_description",
	}
)

func GetTransactionFilePath() string {
	return path.Join(TRANSACATION_FOLDER_PATH, "transactions.csv")
}

func CreateEmptyTransactionFile() (*TransactionFile, *table.Storage, error) {
	logs := NewTransactionFile()
	storage := table.NewStorage()
	err := logs.Save()
	return logs, storage, err
}

func AddBlock(block []events.IEvent, storage *table.Storage, transactionFile *TransactionFile,
	prevTransactionName string, prevTransactionTimeUnix int,
) (*table.Storage, *TransactionFile, error) {

	if len(block) == 1 {
		_, ok := block[0].(*events.WriteEvent)
		if ok {
			transactionFile.addTransaction(block, prevTransactionName, prevTransactionTimeUnix, false)
			return storage, transactionFile, nil
		}
	}

	trans, err := FromEvents(block)
	if err != nil {
		return nil, nil, err
	}

	st, err := trans.Eval(*storage, transactionFile, prevTransactionName, prevTransactionTimeUnix, false)
	if err != nil {
		return nil, nil, err
	}
	storage = &st

	// transactionFile.addTransaction(block, prevTransactionName, prevTransactionTimeUnix, false)

	return storage, transactionFile, nil

}

func LoadTransactionFile() (*TransactionFile, *table.Storage, error) {
	file, err := os.OpenFile(GetTransactionFilePath(), os.O_RDONLY, 0600)
	if err != nil {
		return CreateEmptyTransactionFile()
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		return nil, nil, err
	}

	for i := range columns {
		if columns[i] != records[0][i] {
			return nil, nil, errors.New("error in transaction file - columns doesn't match")
		}
	}

	transactionFile := NewTransactionFile()
	storage := table.NewStorage()

	lastCopmlexTransactionEvents := make([]events.IEvent, 0)
	prevTransactionName := ""
	prevTransactionTimeUnix := -1

	for _, line := range records[1:] {
		transactionTimeUnix, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, nil, errors.New("error in transaction file")
		}
		// transactionTime := time.Unix(0, int64(transactionTimeUnix))
		transactionName := line[1]

		eventType := line[2]
		// eventTable := line[3]
		eventDescription := line[3]

		// event := &events.Event{TableName: eventTable}
		event := &events.Event{}
		// isComplex := false

		if eventType == string(events.CreateEventType) {
			Ev := &events.CreateEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.DeleteEventType) {
			Ev := &events.DeleteEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.DropEventType) {
			Ev := &events.DropEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.InsertEventType) {
			Ev := &events.InsertEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.RollbackEventType) {
			Ev := &events.RollbackEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.UpdateEventType) {
			Ev := &events.UpdateEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else if eventType == string(events.WriteEventType) {
			Ev := &events.WriteEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), Ev)

			event.IEvent = Ev
		} else {
			return nil, nil, errors.New("unknown event type")
		}

		// if isComplex {
		if len(lastCopmlexTransactionEvents) > 0 && prevTransactionName != transactionName {
			// execs, err := executor.FromEvents(lastCopmlexTransactionEvents)

			// trans := NewComplexTransaction(execs)
			st, logs, err := AddBlock(lastCopmlexTransactionEvents, storage, transactionFile, prevTransactionName, prevTransactionTimeUnix)
			if err != nil {
				return nil, nil, err
			}
			storage = st
			transactionFile = logs

			lastCopmlexTransactionEvents = make([]events.IEvent, 0)
		}

		lastCopmlexTransactionEvents = append(lastCopmlexTransactionEvents, event.IEvent)
		// } else {

		// }

		// log := &Log{
		// 	TransactionTime: transactionTime,
		// 	TransactionName: transactionName,
		// 	Ev:              event.IEvent,
		// }

		// transactionFile.Logs = append(transactionFile.Logs, log)
		prevTransactionName = transactionName
		prevTransactionTimeUnix = transactionTimeUnix

	}

	if len(lastCopmlexTransactionEvents) > 0 {
		st, logs, err := AddBlock(lastCopmlexTransactionEvents, storage, transactionFile, prevTransactionName, prevTransactionTimeUnix)
		if err != nil {
			return nil, nil, err
		}
		storage = st
		transactionFile = logs
	}

	return transactionFile, storage, nil
}

func (logs *TransactionFile) Save() error {
	err := os.MkdirAll(TRANSACATION_FOLDER_PATH, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(GetTransactionFilePath(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	data := [][]string{columns}

	for _, log := range logs.Logs {
		record := []string{
			strconv.Itoa(int(log.TransactionTime.UnixNano())),
			log.TransactionName,
			log.GetEventType(),
			// log.GetTableName(),
			log.Ev.GetDescription(),
		}

		data = append(data, record)
	}

	w := csv.NewWriter(file)
	w.Comma = ','
	w.WriteAll(data)
	return w.Error()
}
