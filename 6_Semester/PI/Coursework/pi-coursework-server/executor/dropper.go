package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
)

type Dropper struct {
	TableName string
}

func NewDropper(tableName string) *Dropper {
	return &Dropper{
		TableName: tableName,
	}
}

func (dropper *Dropper) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()

	tab, err := copied.GetTable(dropper.TableName)
	if err != nil {
		return *storage, nil, err
	}

	event := events.NewDropEvent(tab.TableName, tab.Columns, tab.Values)

	err = copied.DropTable(dropper.TableName)

	return *copied, event, err
}
