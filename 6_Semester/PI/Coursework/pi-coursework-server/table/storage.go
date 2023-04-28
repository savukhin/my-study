package table

import "errors"

type Storage struct {
	tables map[string]*Table
}

func NewStorage() *Storage {
	return &Storage{
		tables: map[string]*Table{},
	}
}

func (storage *Storage) GetTables() map[string]*Table {
	return storage.tables
}

func (storage *Storage) AddTable(new_table *Table) {
	copied := new(Table)

	*copied = *new_table
	storage.tables[copied.TableName] = copied
}

func (storage *Storage) GetTableCopy(name string) (Table, error) {
	val, ok := storage.tables[name]
	if !ok {
		return Table{}, errors.New("no such table " + name)
	}

	return *val, nil
}
