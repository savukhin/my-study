package planner

const SelectorName = "Selector"

type Selector struct {
	Table string
	IProcessor
}

func NewSelector(table string) *Selector {
	return &Selector{
		Table: table,
	}
}

func (selector *Selector) DoProcess() *Table {
	// Table
	return nil
}
