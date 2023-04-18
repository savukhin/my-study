package processors

type Aggregator struct {
	Column string
	Sign   string
	Value  string
	IProcessor
}

func NewAggregator(column string, sign string, value string) *Aggregator {
	return &Aggregator{
		Column: column,
		Sign:   sign,
		Value:  value,
	}
}
