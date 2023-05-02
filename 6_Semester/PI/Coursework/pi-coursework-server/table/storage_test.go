package table

import (
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
