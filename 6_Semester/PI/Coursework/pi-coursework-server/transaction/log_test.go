package transaction

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"pi-coursework-server/events"
	"pi-coursework-server/executor"
	"pi-coursework-server/table"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	// ------ FIXTURE LIKE ------ //
	ex, err := os.Executable()
	require.NoError(t, err)

	exPath := filepath.Dir(ex)

	TRANSACATION_FILE_PATH = path.Join(exPath, "transactions.csv")

	{
		// Ev := &CreateEvent{}
		// var down IEvent
		// down = Ev
		// _, ok := down.(*DeleteEvent)
		// fmt.Println("REAL OK")
		// require.True(t, ok)
	}

	t.Log("Write&Load transaction log")
	{
		logs := NewTransactionFile()
		logs.AddCreateEvent(
			[]string{"username", "password"},
			[][]string{
				{"Petr", "petya_king_2014"},
				{"Kirill", "king_of_gorril"},
				{"Masha", "mariya_toropova"},
			},
			"Users",
			"",
			-1,
			true,
		)

		logs.AddCreateEvent(
			[]string{"username2", "passowrd"},
			[][]string{
				{"Named stuff", "Real_slim_shady"},
			},
			"Users2",
			"Transaction1",
			-1,
			true,
		)

		// logs.AddDeleteEvent(
		// 	[]int{3, 10, 1},
		// 	"UsersDelete",
		// 	"Delete_transaction",
		// 	-1,
		// 	true,
		// )

		// err := logs.Save()
		// require.NoError(t, err)

		fmt.Println("0", logs.Logs[0].TransactionName)
		fmt.Println("1", logs.Logs[1].TransactionName)

		q, _ := ioutil.ReadFile(TRANSACATION_FILE_PATH)
		print(string(q))

		logs_loaded, _, err := LoadTransactionFile()
		fmt.Println("0", logs_loaded.Logs[0].TransactionName)
		fmt.Println("1", logs_loaded.Logs[1].TransactionName)
		fmt.Println("logs", logs_loaded.Logs)
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))

		for i := range logs_loaded.Logs {
			require.Equal(t, logs_loaded.Logs[i].TransactionName, logs.Logs[i].TransactionName)
			// require.EqualValues(t, logs_loaded.Logs[i].TransactionTime.UnixNano(), logs.Logs[i].TransactionTime.UnixNano())

			if logs.Logs[i].Ev.GetEventType() == string(events.CreateEventType) {
				event_loaded, ok := logs_loaded.Logs[i].Ev.(*events.CreateEvent)
				require.True(t, ok)
				event_local, ok := logs.Logs[i].Ev.(*events.CreateEvent)
				require.True(t, ok)

				require.Equal(t, event_loaded, event_local)
				require.Equal(t, event_loaded.TableName, event_local.TableName)
				// require.Equal(t, event_loaded.Columns, event_local.Columns)
				// require.Equal(t, event_loaded.Lines, event_local.Lines)
			} else if logs.Logs[i].Ev.GetEventType() == string(events.DeleteEventType) {
				event_local, ok := logs.Logs[i].Ev.(*events.DeleteEvent)
				require.True(t, ok)
				event_loaded, ok := logs_loaded.Logs[i].Ev.(*events.DeleteEvent)
				require.True(t, ok)

				require.Equal(t, event_loaded, event_local)
				require.Equal(t, event_loaded.TableName, event_local.TableName)
				require.Equal(t, event_loaded.Indexes, event_local.Indexes)
			} else {
				require.NoError(t, errors.New("event type not recognized"))
			}
		}
	}

	t.Log("Complex transaction")
	{
		storage := table.NewStorage()

		logs := NewTransactionFile()
		complex := NewComplexTransaction(
			[]executor.IExecutor{
				executor.IExecutor(executor.NewCreator("users", []string{"username", "password"})),
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Mike", "Shinoda"})),
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Chester", "Bennington"})),
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Bob", "Dvlan"})),
				executor.IExecutor(executor.MustNewUpdater("users", "password", "==", "Dvlan", map[string]string{"password": "Dylan"})),
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Iggy", "Pop"})),
			},
		)

		storage2, err := complex.Eval(*storage, logs, "", -1, true)
		require.NoError(t, err)
		users_table, err := storage2.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, []string{"username", "password"}, users_table.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
			{"Bob", "Dylan"},
			{"Iggy", "Pop"},
		}, users_table.Values)

		err = logs.Save()
		require.NoError(t, err)

		logs_loaded, _, err := LoadTransactionFile()
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))
	}

}

func TestPipeline(t *testing.T) {
	// ------ FIXTURE LIKE ------ //
	ex, err := os.Executable()
	require.NoError(t, err)

	exPath := filepath.Dir(ex)

	TRANSACATION_FILE_PATH = path.Join(exPath, "transactions.csv")
	table.TABLES_PATH = path.Join(exPath, "tables")

	storage := *table.NewStorage()

	logs := NewTransactionFile()

	t.Log("Complex transactions")
	{
		storage, err = NewComplexTransaction(
			[]executor.IExecutor{
				executor.IExecutor(executor.NewCreator("users", []string{"username", "password"})),
			},
		).Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		storage, err = NewComplexTransaction(
			[]executor.IExecutor{
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Mike", "Shinoda"})),
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Chester", "Bennington"})),
			},
		).Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		storage, err = NewComplexTransaction(
			[]executor.IExecutor{
				executor.IExecutor(executor.MustNewInserter("users", []string{"username", "password"}, []string{"Bob", "Dvlan"})),
			},
		).Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		storage, err = NewComplexTransaction(
			[]executor.IExecutor{
				executor.IExecutor(executor.MustNewUpdater("users", "password", "==", "Dvlan", map[string]string{"password": "Dylan"})),
			},
		).Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		users_table, err := storage.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, []string{"username", "password"}, users_table.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
			{"Bob", "Dylan"},
		}, users_table.Values)

		err = logs.Save()
		require.NoError(t, err)

		require.Equal(t, 5, len(logs.Logs))

		logs_loaded, _, err := LoadTransactionFile()
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))
	}

	t.Log("Write transactions 1")
	{
		storage, err = NewWriteTransaction().Eval(storage, logs)
		require.NoError(t, err)

		require.Equal(t, 6, len(logs.Logs))

		storage2, err := table.LoadStorage()
		require.NoError(t, err)

		users, err := storage2.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, "users", users.TableName)
		require.Equal(t, []string{"username", "password"}, users.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
			{"Bob", "Dylan"},
		}, users.Values)
	}

	t.Log("Rollback transactions 1")
	{
		storage, err = NewRollbackTransaction().Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		require.Equal(t, 7, len(logs.Logs))

		users_table, err := storage.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, []string{"username", "password"}, users_table.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
			{"Bob", "Dvlan"},
		}, users_table.Values)

		// Repeat

		storage, err = NewRollbackTransaction().Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		require.Equal(t, 8, len(logs.Logs))

		users_table, err = storage.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, []string{"username", "password"}, users_table.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
		}, users_table.Values)
	}

	t.Log("Write transactions 2")
	{
		storage, err = NewWriteTransaction().Eval(storage, logs)
		require.NoError(t, err)

		require.Equal(t, 9, len(logs.Logs))

		storage2, err := table.LoadStorage()
		require.NoError(t, err)

		users, err := storage2.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, "users", users.TableName)
		require.Equal(t, []string{"username", "password"}, users.Columns)
		require.Equal(t, [][]string{
			{"Mike", "Shinoda"},
			{"Chester", "Bennington"},
		}, users.Values)
	}

	t.Log("Rollback transactions 2")
	{
		storage, err = NewRollbackTransaction().Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		require.Equal(t, 10, len(logs.Logs))

		users_table, err := storage.GetTable("users")
		require.NoError(t, err)

		require.Equal(t, []string{"username", "password"}, users_table.Columns)
		require.Equal(t, [][]string{}, users_table.Values)

		// err = logs.Save()
		// require.NoError(t, err)

		q, _ := ioutil.ReadFile(TRANSACATION_FILE_PATH)
		fmt.Println("TRANSACTION FILE", string(q))

		logs_loaded, _, err := LoadTransactionFile()
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))
	}

	t.Log("Rest rollback transactions (clear)")
	{
		// Repeat -- this is clear

		storage, err = NewRollbackTransaction().Eval(storage, logs, "", -1, true)
		require.NoError(t, err)

		require.Equal(t, 11, len(logs.Logs))

		_, err := storage.GetTable("users")
		require.Error(t, err)

		err = logs.Save()
		require.NoError(t, err)

		logs_loaded, _, err := LoadTransactionFile()
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))
	}

	t.Log("Rest write transactions")
	{
		storage, err = NewWriteTransaction().Eval(storage, logs)
		require.NoError(t, err)

		require.Equal(t, 12, len(logs.Logs))

		storage2, err := table.LoadStorage()
		require.NoError(t, err)

		require.Equal(t, 0, len(storage2.GetTables()))
	}
}
