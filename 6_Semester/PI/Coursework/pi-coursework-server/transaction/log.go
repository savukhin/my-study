package transaction

import (
	"pi-coursework-server/events"
	"time"
)

type Log struct {
	TransactionTime time.Time
	TransactionName string
	Ev              events.IEvent
}

func (log *Log) GetEventType() string {
	return log.Ev.GetEventType()
	// stf, ok := log.Ev.(*CreateEvent)
	// fmt.Println("CREATE EVENT???", ok, log.Ev, stf)
	// _, del_ok := log.Ev.(*DeleteEvent)
	// fmt.Println("DELETE EVENT???", del_ok)
	// // _, ev_ok := log.Ev.(Event)
	// // fmt.Println("EVENT EVENT???", ev_ok)
	// if ok {
	// 	return CreateEventType
	// }
	// return DeleteventType
}

func (log *Log) GetTableName() string {
	return log.Ev.GetTableName()
}
