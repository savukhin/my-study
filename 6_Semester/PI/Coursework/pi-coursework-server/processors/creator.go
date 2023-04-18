package processors

type TableCreator struct {
	Table   string
	Columns []string
	IProcessor
}

func NewTableCreator(tableName string, colums []string) *TableCreator {
	return &TableCreator{
		Table:   tableName,
		Columns: colums,
	}
}
