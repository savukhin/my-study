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
	copied := new_table.Copy()

	storage.tables[copied.TableName] = copied
}

func (storage *Storage) MustGetTableCopy(name string) Table {
	val, err := storage.GetTableCopy(name)
	if err != nil {
		panic(err)
	}
	return val
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
