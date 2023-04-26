package table

type Dimensions struct {
	X int
	Y int
}

type IndexRange struct {
	Start int32
	End   int32
}

type Table struct {
	TableName string
	Columns   []string
	Values    [][]string
	Shape     Dimensions
}

func NewTable() *Table {
	return &Table{}
}
