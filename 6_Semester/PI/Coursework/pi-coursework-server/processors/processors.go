package processors

import "pi-coursework-server/table"

type IProcessor interface {
	DoProcess(table.Storage) (table.Storage, error)
}
