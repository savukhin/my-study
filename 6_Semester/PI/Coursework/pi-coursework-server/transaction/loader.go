package transaction

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
	"path"
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

func LoadTransactionFile() (*TransactionFile, error) {
	file, err := os.OpenFile(TRANSACATION_FILE_PATH, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

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

		event := &Event{TableName: eventTable}

		if eventType == string(CreateEventType) {
			ev := &CreateEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), ev)

			event.IEvent = ev
		} else {
			ev := &DeleteEvent{Event: event}
			json.Unmarshal([]byte(eventDescription), ev)

			event.IEvent = ev
		}

		log := &Log{
			TransactionTime: transactionTime,
			TransactionName: transactionName,
			ev:              event.IEvent,
		}

		transactionFile.Logs = append(transactionFile.Logs, log)

	}

	return transactionFile, nil
}

func (logs *TransactionFile) Save() error {
	file, err := os.OpenFile(TRANSACATION_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0600)
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
			log.ev.GetDescription(),
		}

		data = append(data, record)
	}

	w := csv.NewWriter(file)
	w.Comma = ','
	w.WriteAll(data)
	return w.Error()
}
