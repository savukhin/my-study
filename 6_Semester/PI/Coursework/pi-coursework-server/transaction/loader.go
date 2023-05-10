package transaction

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
	"path"
	"pi-coursework-server/events"
	"pi-coursework-server/utils"
	"strconv"
	"time"
)

var (
	TRANSACATION_FILE_PATH = utils.GetEnvDefault("TRANSACATION_FILE_PATH", path.Join(".", "static", "transactions", "transactions.csv"))
	columns                = []string{
		"time",
		"transaction_name",
		"event_type",
		"event_table",
		"event_description",
	}

	columnsMap = map[string]int{
		"time":              0,
		"transaction_name":  1,
		"event_type":        2,
		"event_table":       3,
		"event_description": 4,
	}
)

func CreateEmptyTransactionFile() (*TransactionFile, error) {
	logs := NewTransactionFile()
	err := logs.Save()
	return logs, err
}

func LoadTransactionFile() (*TransactionFile, error) {
	file, err := os.OpenFile(TRANSACATION_FILE_PATH, os.O_RDONLY, 0600)
	if err != nil {
		return CreateEmptyTransactionFile()
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	for i := range columns {
		if columns[i] != records[0][i] {
			return nil, errors.New("error in transaction file - columns doesn't match")
		}
	}

	transactionFile := NewTransactionFile()

	for _, line := range records[1:] {
		transactionTimeUnix, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, errors.New("error in transaction file")
		}
		transactionTime := time.Unix(0, int64(transactionTimeUnix))
		transactionName := line[1]

		eventType := line[2]
		eventTable := line[3]
		eventDescription := line[4]

		event := &events.Event{TableName: eventTable}

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
			return nil, errors.New("unknown event type")
		}

		log := &Log{
			TransactionTime: transactionTime,
			TransactionName: transactionName,
			Ev:              event.IEvent,
		}

		transactionFile.Logs = append(transactionFile.Logs, log)

	}

	return transactionFile, nil
}

func (logs *TransactionFile) Save() error {
	file, err := os.OpenFile(TRANSACATION_FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
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
			log.GetTableName(),
			log.Ev.GetDescription(),
		}

		data = append(data, record)
	}

	w := csv.NewWriter(file)
	w.Comma = ','
	w.WriteAll(data)
	return w.Error()
}
