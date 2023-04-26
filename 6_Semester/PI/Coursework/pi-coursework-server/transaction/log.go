package transaction

import (
	"time"
)

type Log struct {
	TransactionTime time.Time
	TransactionName string
	ev              IEvent
}

func (log *Log) GetEventType() string {
	return log.ev.GetEventType()
	// stf, ok := log.ev.(*CreateEvent)
	// fmt.Println("CREATE EVENT???", ok, log.ev, stf)
	// _, del_ok := log.ev.(*DeleteEvent)
	// fmt.Println("DELETE EVENT???", del_ok)
	// // _, ev_ok := log.ev.(Event)
	// // fmt.Println("EVENT EVENT???", ev_ok)
	// if ok {
	// 	return CreateEventType
	// }
	// return DeleteventType
}

func (log *Log) GetTableName() string {
	return log.ev.GetTableName()
}
