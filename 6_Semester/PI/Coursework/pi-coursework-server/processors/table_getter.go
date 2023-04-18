package processors

type TableGetter struct {
	Table string
	IProcessor
}

func NewTableGetter(table string) *TableGetter {
	return &TableGetter{
		Table: table,
	}
}
