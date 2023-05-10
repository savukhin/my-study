package executor

import (
	"pi-coursework-server/events"
	"pi-coursework-server/table"
	"pi-coursework-server/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutor(t *testing.T) {
	exPath := utils.GetExecutablePath()
	table.TABLES_PATH = exPath

	users_table := table.MustNewTable("users",
		[]string{"username", "password"},
		[][]string{
			{"Petrarh", "Petya_2001"},
			{"Vitek", "cool_guy_2014"},
		},
	)

	rooms_table := table.MustNewTable("rooms",
		[]string{"floor", "area"},
		[][]string{
			{"15", "12"},
			{"15", "10"},
			{"15", "20"},
			{"12", "10"},
			{"12", "15"},
		},
	)

	t.Log("Selector test")
	{

		storage := table.NewStorage()

		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		require.Equal(t, 2, len(storage.GetTables()))

		selector := NewSelector(
			"rooms",
			[]string{"area"},
			utils.NewWhereConditionCheck("floor", "==", "15"),
			nil,
		)

		table, err := selector.DoExecute(storage)
		require.NoError(t, err)
		require.Equal(t, "rooms", table.TableName)
		require.Equal(t, []string{"area"}, table.Columns)
		require.Equal(t, [][]string{{"12"}, {"10"}, {"20"}}, table.Values)

		selector = NewSelector(
			"rooms",
			[]string{"floor"},
			utils.NewWhereConditionCheck("area", "!=", "12"),
			utils.NewLimitCondition(3),
		)

		table, err = selector.DoExecute(storage)
		require.NoError(t, err)
		require.Equal(t, "rooms", table.TableName)
		require.Equal(t, []string{"floor"}, table.Columns)
		require.Equal(t, [][]string{{"15"}, {"15"}, {"12"}}, table.Values)
	}

	t.Log("Creator test")
	{
		storage := table.NewStorage()
		creator := NewCreator("users", []string{"username", "password"})
		require.Equal(t, 0, len(storage.GetTables()))

		storage2, event, err := creator.DoExecute(storage)
		require.NoError(t, err)
		// require.Equal(t, table.Columns, []string{"username", "password"})
		require.Equal(t, 1, len(storage2.GetTables()))

		createEvent, ok := event.(*events.CreateEvent)
		require.True(t, ok)
		require.Equal(t, "users", createEvent.TableName)
	}

	t.Log("Dropper test")
	{
		storage := table.NewStorage()
		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		dropper := NewDropper("users")
		require.Equal(t, 2, len(storage.GetTables()))

		storage2, event, err := dropper.DoExecute(storage)

		require.NoError(t, err)
		require.Equal(t, 1, len(storage2.GetTables()))
		require.Equal(t, 2, len(storage.GetTables()))

		dropEvent, ok := event.(*events.DropEvent)
		require.True(t, ok)
		require.Equal(t, "users", dropEvent.TableName)
	}

	t.Log("Inserter test")
	{
		storage := table.NewStorage()
		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		inserter := NewInserterFromMap("users", map[string]string{
			"username": "Les",
			"password": "Paul",
		})
		require.Equal(t, 2, len(storage.GetTables()))

		storage2, event, err := inserter.DoExecute(storage)

		require.NoError(t, err)
		require.Equal(t, 2, storage.MustGetTable("users").Shape.Y)
		require.Equal(t, 3, storage2.MustGetTable("users").Shape.Y)

		insertEvent, ok := event.(*events.InsertEvent)
		require.True(t, ok)
		require.Equal(t, "users", insertEvent.TableName)
		require.Equal(t, []string{"Les", "Paul"}, insertEvent.Values)

		_, _, err = NewInserterFromMap("users", map[string]string{
			"username":  "Les",
			"password":  "Paul",
			"excessive": "Gibson",
		}).DoExecute(storage)

		require.Error(t, err)
	}

	t.Log("Updater test")
	{
		storage := table.NewStorage()
		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		updater, _ := NewUpdater("users", "username", "==", "Vitek", map[string]string{
			"username": "Bob",
		})
		require.Equal(t, 2, len(storage.GetTables()))

		storage2, event, err := updater.DoExecute(storage)

		require.NoError(t, err)

		require.Equal(t, "Vitek", storage.MustGetTable("users").Values[1][0])
		require.Equal(t, "Bob", storage2.MustGetTable("users").Values[1][0])

		updateEvent, ok := event.(*events.UpdateEvent)
		require.True(t, ok)
		require.Equal(t, []int{1}, updateEvent.Indexes)
		require.Equal(t, map[int][]string{1: []string{"Bob", "cool_guy_2014"}}, updateEvent.Values)

	}

	t.Log("Deleter test")
	{
		storage := table.NewStorage()
		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		deleter := NewDeleter("rooms", "floor", "==", "15")
		require.Equal(t, 2, len(storage.GetTables()))

		require.Equal(t, 5, storage.MustGetTable("rooms").Shape.Y)

		storage2, event, err := deleter.DoExecute(storage)

		require.NoError(t, err)
		require.Equal(t, 5, storage.MustGetTable("rooms").Shape.Y)
		require.Equal(t, 2, storage2.MustGetTable("rooms").Shape.Y)

		deleteEvent, ok := event.(*events.DeleteEvent)
		require.True(t, ok)
		require.Equal(t, []int{0, 1, 2}, deleteEvent.Indexes)
		require.Equal(t, "rooms", deleteEvent.TableName)
	}
}

func TestRollbackEvent(t *testing.T) {
	// createEvent := events.NewCreateEvent("users")

}
