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
		return NewMassCreator(dropEvent.TableName, dropEvent.Columns, dropEvent.Values), nil
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

	return nil, errors.New("no such rollback executor for this event")
}

func FromEvent(event events.IEvent) (IExecutor, error) {
	createEvent, ok := event.(*events.CreateEvent)
	if ok {
		return NewCreator(createEvent.TableName, createEvent.Columns), nil
	}

	deleteEvent, ok := event.(*events.DeleteEvent)
	if ok {
		massDeleter := NewMassDeleter(deleteEvent.TableName, deleteEvent.Indexes)
		return massDeleter, nil
	}

	dropEvent, ok := event.(*events.DropEvent)
	if ok {
		return NewDropper(dropEvent.TableName), nil
	}

	insertEvent, ok := event.(*events.InsertEvent)
	if ok {
		values := map[int][]string{
			insertEvent.Index: insertEvent.Values,
		}

		// return NewInserter(insertEvent.TableName, insertEvent.Index, insertEvent.Values)
		return NewMassInserter(insertEvent.TableName, []int{insertEvent.Index}, values)
	}

	updateEvent, ok := event.(*events.UpdateEvent)
	if ok {
		updater := NewMassUpdater(updateEvent.TableName, updateEvent.Indexes, updateEvent.Values)

		return updater, nil
	}

	return nil, errors.New("no such executor for this event")
}

func FromEvents(eventsArr []events.IEvent) ([]IExecutor, error) {
	result := make([]IExecutor, len(eventsArr))
	for i, ev := range eventsArr {
		exec, err := FromEvent(ev)
		if err != nil {
			return nil, err
		}
		result[i] = exec
	}
	return result, nil
}
