package table

import (
	"encoding/csv"
	"errors"
	"os"
	"path"
	"pi-coursework-server/utils"
)

var (
	TABLES_PATH = utils.GetEnvDefault("TABLES_PATH", path.Join(".", "static", "tables"))
)

func LoadTable(fileNameCsv string) (*Table, error) {
	if len(fileNameCsv) <= 4 {
		return nil, errors.New("small filename")
	}

	tableName := fileNameCsv[:len(fileNameCsv)-4]

	if fileNameCsv[len(tableName):] != ".csv" {
		return nil, errors.New("not a csv")
	}

	filePath := path.Join(TABLES_PATH, fileNameCsv)
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, errors.New("no data")
	}

	table := &Table{
		TableName: tableName,
		Columns:   records[0],
		Values:    records[1:],
		Shape: Dimensions{
			X: len(records[0]),
			Y: len(records) - 1,
		},
	}

	return table, nil
}

func LoadAllTables() ([]*Table, error) {
	dir, err := os.ReadDir(TABLES_PATH)
	if err != nil {
		return nil, err
	}

	tables := make([]*Table, 0)

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		table, err := LoadTable(file.Name())
		if err != nil {
			continue
		}

		tables = append(tables, table)
	}

	return tables, nil
}
