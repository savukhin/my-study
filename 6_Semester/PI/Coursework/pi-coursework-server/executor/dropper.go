package executor

import (
	"pi-coursework-server/table"
)

type Dropper struct {
	IExecutor

	TableName string
}

func NewDropper(tableName string) *Dropper {
	return &Dropper{
		TableName: tableName,
	}
}

func (dropper *Dropper) DoExecute(storage *table.Storage) (table.Storage, error) {
	copied := storage.Copy()

	err := copied.DropTable(dropper.TableName)

	return *copied, err
}
