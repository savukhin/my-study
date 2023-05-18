package table

import (
	"pi-coursework-server/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Log("Checking if Copy() works")
	{
		users_table := MustNewTable("users",
			[]string{"username", "password"},
			[][]string{
				{"Petrarh", "Petya_2001"},
				{"Vitek", "cool_guy_2014"},
			},
		)

		rooms_table := MustNewTable("rooms",
			[]string{"floor", "area"},
			[][]string{
				{"15", "12"},
				{"15", "10"},
				{"15", "20"},
				{"12", "10"},
				{"12", "15"},
			},
		)

		storage := NewStorage()
		storage.AddTable(rooms_table)

		storage2 := storage.Copy()
		storage2.AddTable(users_table)

		require.Equal(t, 1, len(storage.tables))
		require.Equal(t, 2, len(storage2.tables))

		require.Equal(t, 5, storage.MustGetTable("rooms").Shape.Y)
		require.Equal(t, 5, storage2.MustGetTable("rooms").Shape.Y)
		// rooms_table.AddRow([]string{"1", "1"})

		storage.MustGetTable("rooms").AddRow([]string{"1", "1"})

		require.Equal(t, 6, storage.MustGetTable("rooms").Shape.Y)
		require.Equal(t, 5, storage2.MustGetTable("rooms").Shape.Y)

		storage.AddTable(users_table)

		storage2.MustGetTable("users").Values[0][0] = "Alice"
		require.Equal(t, "Petrarh", storage.MustGetTable("users").Values[0][0])
		require.Equal(t, "Alice", storage2.MustGetTable("users").Values[0][0])
	}
}

func TestSaveAndLoad(t *testing.T) {
	// ------ FIXTURE LIKE ------ //
	exPath := utils.GetExecutablePath()

	TABLES_PATH = exPath

	// os.RemoveAll(path.Join(TABLES_PATH, "/"))

	t.Log("Save&Load test")
	{
		users, err := NewTable("users", []string{"first_name", "last_name", "username"}, [][]string{
			{"Rob", "Pike", "rob"},
			{"Ken", "Thompson", "ken"},
			{"Robert", "Griesemer", "gri"},
		})
		require.NoError(t, err)

		rooms, err := NewTable("rooms", []string{"room_size", "room_number", "room_floor", "owner"}, [][]string{
			{"1", "14", "3", "Greg"},
			{"2", "13", "1", "Andrew"},
			{"3", "12", "15", "Fahren"},
			{"4", "10", "2", "Igor"},
			{"3", "15", "3", "Jora"},
		})
		require.NoError(t, err)

		storage := NewStorage()
		storage.AddTable(users)
		storage.AddTable(rooms)
		storage.Save()

		storageLoaded, err := LoadStorage()
		require.NoError(t, err)

		require.Equal(t, 2, len(storage.tables))
		require.Equal(t, 2, len(storageLoaded.tables))

		usersLoaded, err := storageLoaded.GetTable("users")
		require.NoError(t, err)
		roomsLoaded, err := storageLoaded.GetTable("rooms")
		require.NoError(t, err)

		t.Log("Check shapes")
		require.EqualValues(t, 3, usersLoaded.Shape.X)
		require.EqualValues(t, 3, usersLoaded.Shape.Y)
		require.EqualValues(t, 4, roomsLoaded.Shape.X)
		require.EqualValues(t, 5, roomsLoaded.Shape.Y)

		t.Log("Check columns&values")
		require.Equal(t, users.Columns, usersLoaded.Columns)
		require.Equal(t, users.Values, usersLoaded.Values)
		require.Equal(t, rooms.Columns, roomsLoaded.Columns)
		require.Equal(t, rooms.Values, roomsLoaded.Values)
	}
}
