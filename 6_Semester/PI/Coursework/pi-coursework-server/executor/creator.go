package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type Creator struct {
	TableName string
	Columns   []string
}

func NewCreator(tableName string, columns []string) *Creator {
	creator := &Creator{
		TableName: tableName,
		Columns:   columns,
	}

	// creator.IExecutor = creator
	return creator
}

func (creator *Creator) DoExecute(storage *table.Storage) (table.Storage, error) {
	table := table.MustNewTable(creator.TableName, creator.Columns, make([][]string, 0))
	err := table.Save()
	storage.AddTable(table)
	return *storage, err
	// return *table, err
}

func (creator *Creator) ToEvent() *events.Event {
	// return &events.CreateEvent{
	// 	TableName: creator.TableName,
	// }
	return nil
}
