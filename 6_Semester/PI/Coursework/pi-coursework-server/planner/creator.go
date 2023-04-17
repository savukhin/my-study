package planner

const TableCreatorName = "Table Creator"

type TableCreator struct {
	Table   string
	columns []string
	IProcessor
}

func NewTableCreator(tableName string, colums []string) *TableCreator {
	return &TableCreator{
		Table:   tableName,
		columns: colums,
	}
}

func (creator *TableCreator) GetName() string {
	return TableCreatorName
}
