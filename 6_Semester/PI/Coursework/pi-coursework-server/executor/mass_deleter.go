package executor

import (
	"errors"
	"pi-coursework-server/events"
	"pi-coursework-server/table"
	"sort"
)

type MassDeleter struct {
	TableName string
	Indexes   []int
}

func NewMassDeleter(tableName string, indexes []int) *MassDeleter {
	return &MassDeleter{
		TableName: tableName,
		Indexes:   indexes,
	}
}

func (massDeleter *MassDeleter) DoExecute(storage *table.Storage) (table.Storage, events.IEvent, error) {
	copied := storage.Copy()
	tab, err := copied.GetTable(massDeleter.TableName)
	if err != nil {
		return *copied, nil, err
	}

	yAbsolute := 0
	initialShapeY := tab.Shape.Y
	i := 0
	indexes := massDeleter.Indexes
	sort.Ints(indexes)
	y := 0
	deletedValues := make(map[int][]string, 0)

	for yAbsolute < initialShapeY && yAbsolute < tab.Shape.Y && i < len(indexes) {
		if yAbsolute == indexes[i] {
			row := make([]string, tab.Shape.X)
			copy(row, tab.Values[y])
			deletedValues[y] = row

			err := tab.DeleteRow(y)

			if err != nil {
				return *copied, nil, err
			}

			i++
		} else {
			y++
		}
		yAbsolute++
	}

	if i != len(indexes) {
		return *copied, nil, errors.New("data corrupted")
	}

	return *copied, events.NewDeleteEvent(massDeleter.TableName, indexes, deletedValues), nil
}
