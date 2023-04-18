package processors

type Inserter struct {
	Values map[string]string
	IProcessor
}

func NewInserter(values map[string]string) *Inserter {
	return &Inserter{
		Values: values,
	}
}
