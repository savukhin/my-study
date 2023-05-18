package processors

type Selector struct {
	Columns []string
	IProcessor
}

func NewSelector(columns []string) *Selector {
	return &Selector{
		Columns: columns,
	}
}

// func (selector *Selector) DoProcess() *Table {
// 	// Table
// 	return nil
// }
