package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
	"sort"
)

type MassUpdater struct {
	TableName string
	Indexes   []int
	Values    map[int][]string
}

func NewMassUpdater(tableName string, indexes []int, values map[int][]string) *MassUpdater {
	return &MassUpdater{
		TableName: tableName,
		Values:    values,
		Indexes:   indexes,
	}
}

func (massUpdater *MassUpdater) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(massUpdater.TableName)
	if err != nil {
		return *copied, nil, err
	}

	indexes := massUpdater.Indexes
	sort.Ints(indexes)
	i := 0

	for y, _ := range tab.Values {
		if y == indexes[i] {
			values := massUpdater.Values[y]

			err := tab.HardUpdateRow(y, values)
			if err != nil {
				return *copied, nil, err
			}

			i++
		}
	}
	return *copied, nil, nil
}
