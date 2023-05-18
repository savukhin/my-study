package processors

type Dropper struct {
	Table string
	IProcessor
}

func NewDropper(table string) *Dropper {
	return &Dropper{
		Table: table,
	}
}

// func (selector *Selector) DoProcess() *Table {
// 	// Table
// 	return nil
// }
