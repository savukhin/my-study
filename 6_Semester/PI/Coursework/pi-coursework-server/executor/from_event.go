package executor

import (
	"errors"
	"pi-coursework-server/events"
)

func RollbackEvent(event events.IEvent) (IExecutor, error) {
	createEvent, ok := event.(*events.CreateEvent)
	if ok {
		return NewDropper(createEvent.TableName), nil
	}

	deleteEvent, ok := event.(*events.DeleteEvent)
	if ok {
		massInserter, err := NewMassInserter(deleteEvent.TableName, deleteEvent.Indexes, deleteEvent.DeletedValues)
		return massInserter, err
	}

	dropEvent, ok := event.(*events.DropEvent)
	if ok {
		return NewCreator(dropEvent.TableName, dropEvent.Columns), nil
	}

	insertEvent, ok := event.(*events.InsertEvent)
	if ok {
		return NewMassDeleter(insertEvent.TableName, []int{insertEvent.Index}), nil
	}

	updateEvent, ok := event.(*events.UpdateEvent)
	if ok {
		updater := NewMassUpdater(updateEvent.TableName, updateEvent.Indexes, updateEvent.OldValues)

		return updater, nil
	}

	return nil, errors.New("no such executor for this event")
}
