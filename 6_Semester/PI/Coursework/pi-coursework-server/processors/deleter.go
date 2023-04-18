package processors

type Deleter struct {
	IProcessor
}

func NewDeleter() *Deleter {
	return &Deleter{}
}

// func (selector *Selector) DoProcess() *Table {
// 	// Table
// 	return nil
// }
