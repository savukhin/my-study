package executor

import "pi-coursework-server/table"

type Creator struct {
	IExecutor

	TableName string
	Columns   []string
}

func NewCreator(tableName string, columns []string) *Creator {
	return &Creator{
		TableName: tableName,
		Columns:   columns,
	}
}

func (creator *Creator) DoExecute(storage *table.Storage) (table.Table, error) {
	table := table.MustNewTable(creator.TableName, creator.Columns, make([][]string, 0))
	err := table.Save()
	storage.AddTable(table)
	return *table, err
}
