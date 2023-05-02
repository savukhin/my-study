package table

import (
	"encoding/csv"
	"errors"
	"os"
	"path"
	"strconv"
)

const INDEX_COLUMN_NAME = "index"

type Dimensions struct {
	X int
	Y int
}

type IndexRange struct {
	Start int32
	End   int32
}

type Table struct {
	TableName      string
	Columns        []string
	columnsSet     map[string]int
	indexColumnInd int
	Values         [][]string // [row][col]
	Shape          Dimensions
}

func MustNewTable(tableName string, columns []string, values [][]string) *Table {
	tab, err := NewTable(tableName, columns, values)
	if err != nil {
		panic(err)
	}
	return tab
}

func NewTable(tableName string, columns []string, values [][]string) (*Table, error) {
	tab := &Table{
		TableName: tableName,
		Columns:   columns,
		Values:    values,
		Shape: Dimensions{
			X: len(columns),
			Y: len(values),
		},
		columnsSet: make(map[string]int),
	}

	for y := range values {
		if len(columns) != len(values[y]) {
			return nil, errors.New("values not match columns shape")
		}
	}

	for i, column := range columns {
		tab.columnsSet[column] = i

		if column == INDEX_COLUMN_NAME {
			tab.indexColumnInd = i
		}
	}

	return tab, nil
}

func (table *Table) Save() error {
	filePath := path.Join(TABLES_PATH, table.TableName+".csv")

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	records := [][]string{table.Columns}
	for y := 0; y < table.Shape.Y; y++ {
		line := make([]string, 0)
		for x := 0; x < table.Shape.X; x++ {
			line = append(line, table.Values[x][y])
		}

		records = append(records, line)
	}

	w := csv.NewWriter(file)
	w.Comma = ','
	w.WriteAll(records) // calls Flush internally
	return w.Error()
}

func (table *Table) HasColumn(column string) bool {
	_, ok := table.columnsSet[column]
	return ok
}

func (table *Table) GetColumn(column string) ([]string, error) {
	ind, ok := table.columnsSet[column]
	if !ok {
		return nil, errors.New("no such column " + column)
	}

	col := make([]string, 0)
	for y := 0; y < table.Shape.Y; y++ {
		col = append(col, table.Values[y][ind])
	}

	return col, nil
}

func (table *Table) AddRowMap(values map[string]string) error {
	if len(values) != table.Shape.X {
		return errors.New("values must have len of " + strconv.Itoa(table.Shape.X) +
			" but they have " + strconv.Itoa(len(values)))
	}

	resultValues := make([]string, len(values))
	for i, column := range table.Columns {
		val, ok := values[column]
		if !ok {
			return errors.New("no such column name " + column)
		}
		resultValues[i] = val
	}

	return table.AddRow(resultValues)
}

func (table *Table) AddRow(values []string) error {
	if len(values) != table.Shape.X {
		return errors.New("values must have len of " + strconv.Itoa(table.Shape.X) +
			" but they have " + strconv.Itoa(len(values)))
	}

	table.Values = append(table.Values, values)
	table.Shape.Y++

	return nil
}

func (table *Table) GetRow(y int) ([]string, error) {
	if y < 0 || y >= table.Shape.Y {
		return nil, errors.New("index out of range")
	}

	result := make([]string, 0)
	for i := 0; i < table.Shape.X; i++ {
		// Append from each column
		result = append(result, table.Values[y][i])
	}

	return result, nil
}

func (table *Table) MustGetRow(y int) []string {
	val, err := table.GetRow(y)
	if err != nil {
		panic(err)
	}

	return val
}

func (table *Table) GetSlice(from_y int, to_y int) (Table, error) {
	if from_y < 0 || from_y >= table.Shape.Y || to_y < 0 || to_y >= table.Shape.Y {
		return Table{}, errors.New("index out of range [" + strconv.Itoa(from_y) + ", " + strconv.Itoa(to_y) + "] of shape y = " + strconv.Itoa(table.Shape.Y))
	}

	result, _ := NewTable(table.TableName, table.Columns, make([][]string, 0))
	for y := from_y; y < to_y; y++ {
		result.AddRow(table.MustGetRow(y))
	}

	return *result, nil
}

func (table *Table) Copy() *Table {
	values := make([][]string, len(table.Values))
	columns := make([]string, len(table.Columns))

	copy(values, table.Values)

	copy(columns, table.Columns)

	copied := MustNewTable(table.TableName, columns, values)

	return copied
}
