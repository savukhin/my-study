package processors

type Dimensions struct {
	X int
	Y int
}

type Table struct {
	Elems [][]string
	Shape Dimensions
}
