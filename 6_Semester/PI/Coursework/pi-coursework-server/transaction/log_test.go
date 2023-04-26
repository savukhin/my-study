package transaction

import (
	"errors"
	"os"
	"path"
	"path/filepath"
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
		// ev := &CreateEvent{}
		// var down IEvent
		// down = ev
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
		)

		logs.AddCreateEvent(
			[]string{"username2", "passowrd"},
			[][]string{
				{"Named stuff", "Real_slim_shady"},
			},
			"Users2",
			"Transaction1",
		)

		logs.AddDeleteEvent(
			[]int32{3, 10, 1},
			"UsersDelete",
			"Delete_transaction",
		)

		err := logs.Save()
		require.NoError(t, err)

		logs_loaded, err := LoadTransactionFile()
		require.NoError(t, err)

		require.Equal(t, len(logs_loaded.Logs), len(logs.Logs))

		for i := range logs_loaded.Logs {
			require.Equal(t, logs_loaded.Logs[i].TransactionName, logs.Logs[i].TransactionName)
			require.EqualValues(t, logs_loaded.Logs[i].TransactionTime.UnixNano(), logs.Logs[i].TransactionTime.UnixNano())

			if logs.Logs[i].ev.GetEventType() == string(CreateEventType) {
				event_loaded, ok := logs_loaded.Logs[i].ev.(*CreateEvent)
				require.True(t, ok)
				event_local, ok := logs.Logs[i].ev.(*CreateEvent)
				require.True(t, ok)

				require.Equal(t, event_loaded, event_local)
				require.Equal(t, event_loaded.TableName, event_local.TableName)
				require.Equal(t, event_loaded.Columns, event_local.Columns)
				require.Equal(t, event_loaded.Lines, event_local.Lines)
			} else if logs.Logs[i].ev.GetEventType() == string(DeletEventType) {
				event_local, ok := logs.Logs[i].ev.(*DeleteEvent)
				require.True(t, ok)
				event_loaded, ok := logs_loaded.Logs[i].ev.(*DeleteEvent)
				require.True(t, ok)

				require.Equal(t, event_loaded, event_local)
				require.Equal(t, event_loaded.TableName, event_local.TableName)
				require.Equal(t, event_loaded.Indexes, event_local.Indexes)
			} else {
				require.NoError(t, errors.New("event type not recognized"))
			}
		}

	}
}
