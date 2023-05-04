package table

import (
	"errors"
	"os"
)

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

func (storage *Storage) AddTable(new_table *Table) error {
	_, ok := storage.tables[new_table.TableName]
	if ok {
		return errors.New("table" + new_table.TableName + "already exists")
	}

	copied := new_table.Copy()

	storage.tables[copied.TableName] = copied
	return nil
}

func (storage *Storage) MustGetTableCopy(name string) Table {
	val, err := storage.GetTableCopy(name)
	if err != nil {
		panic(err)
	}
	return val
}

func (storage *Storage) Save() error {
	err := os.RemoveAll(TABLES_PATH)
	if err != nil {
		return err
	}

	os.MkdirAll(TABLES_PATH, os.ModePerm)

	for _, table := range storage.tables {
		err := table.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadStorage() (*Storage, error) {
	tables, err := LoadAllTables()
	if err != nil {
		return nil, err
	}

	storage := NewStorage()

	for _, table := range tables {
		storage.AddTable(table)
	}

	return storage, nil
}

func (storage *Storage) GetTableCopy(name string) (Table, error) {
	val, ok := storage.tables[name]
	if !ok {
		return Table{}, errors.New("no such table " + name)
	}

	return *val, nil
}

func (storage *Storage) MustGetTable(name string) *Table {
	val, err := storage.GetTable(name)
	if err != nil {
		panic(err)
	}
	return val
}

func (storage *Storage) GetTable(name string) (*Table, error) {
	val, ok := storage.tables[name]
	if !ok {
		return nil, errors.New("no such table " + name)
	}

	return val, nil
}

func (storage *Storage) DropTable(name string) error {
	_, ok := storage.tables[name]
	if !ok {
		return errors.New("no such table " + name)
	}

	delete(storage.tables, name)
	return nil
}

func (storage *Storage) Copy() *Storage {
	copied := NewStorage()
	for key, table := range storage.tables {
		copied.tables[key] = table.Copy()
	}

	return copied
}
