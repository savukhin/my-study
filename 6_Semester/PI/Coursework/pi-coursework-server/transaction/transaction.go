package transaction

import "pi-coursework-server/table"

type ITransaction interface {
	Eval(table.Storage, *TransactionFile) error
}

// type AbstractTransaction struct {
// 	ITransaction
// }
