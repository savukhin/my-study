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
	// Elems [][]string

	// Shape Dimensions

	Ranges []IndexRange
}

func NewTable() *Table {
	return &Table{}
}
