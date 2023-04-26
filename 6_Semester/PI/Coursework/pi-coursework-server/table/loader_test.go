package table

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRecords(records [][]string, fileName string) error {
	file, err := os.OpenFile(path.Join(TABLES_PATH, fileName), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	w.Comma = ','
	w.WriteAll(records) // calls Flush internally
	return w.Error()

}

func TestLoader(t *testing.T) {
	// ------ FIXTURE LIKE ------ //
	ex, err := os.Executable()
	require.NoError(t, err)

	exPath := filepath.Dir(ex)

	TABLES_PATH = exPath

	t.Log("Load test")
	{
		users := [][]string{
			{"first_name", "last_name", "username"},
			{"Rob", "Pike", "rob"},
			{"Ken", "Thompson", "ken"},
			{"Robert", "Griesemer", "gri"},
		}

		rooms := [][]string{
			{"room_size", "room_number", "room_floor", "owner"},
			{"1", "14", "3", "Greg"},
			{"2", "13", "1", "Andrew"},
			{"3", "12", "15", "Fahren"},
			{"4", "10", "2", "Igor"},
			{"3", "15", "3", "Jora"},
		}

		require.NoError(t, createRecords(users, "users.csv"))
		require.NoError(t, createRecords(rooms, "rooms.csv"))

		tables, err := LoadAllTables()
		require.NoError(t, err)

		require.Equal(t, 2, len(tables))

		var usersTable *Table
		var roomsTable *Table

		t.Log("Check table names")
		if tables[0].TableName == "users" && tables[1].TableName == "rooms" {
			usersTable = tables[0]
			roomsTable = tables[1]
		} else if tables[1].TableName == "users" && tables[0].TableName == "rooms" {
			usersTable = tables[1]
			roomsTable = tables[0]
		} else {
			fmt.Println(tables[0].TableName, tables[1].TableName)
			require.Truef(t, false, "Not users or rooms tables!")
		}

		t.Log("Check shapes")
		require.EqualValues(t, 3, usersTable.Shape.X)
		require.EqualValues(t, 3, usersTable.Shape.Y)
		require.EqualValues(t, 4, roomsTable.Shape.X)
		require.EqualValues(t, 5, roomsTable.Shape.Y)

		t.Log("Check columns")
		require.Equal(t, users[0], usersTable.Columns)
		require.Equal(t, users[1:], usersTable.Values)
		require.Equal(t, rooms[0], roomsTable.Columns)
		require.Equal(t, rooms[1:], roomsTable.Values)

	}
}
