package executor

import (
	"encoding/csv"
	"os"
	"path"
	"pi-coursework-server/planner"
	"pi-coursework-server/table"
	"pi-coursework-server/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutor(t *testing.T) {
	exPath := utils.GetExecutablePath()

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
		table.TABLES_PATH = exPath

		storage := table.NewStorage()

		storage.AddTable(users_table)
		storage.AddTable(rooms_table)

		require.Equal(t, 2, len(storage.GetTables()))

		selector := NewSelector(
			"rooms",
			planner.NewWhereConditionCheck("floor", "==", "15"),
			nil,
		)

		table, err := selector.DoExecute(storage)
		require.NoError(t, err)
		require.Equal(t, "rooms", table.TableName)
		require.Equal(t, []string{"floor", "area"}, table.Columns)
		require.Equal(t, [][]string{{"15", "12"}, {"15", "10"}, {"15", "20"}}, table.Values)

		selector = NewSelector(
			"rooms",
			planner.NewWhereConditionCheck("area", "!=", "12"),
			planner.NewLimitCondition(3),
		)

		table, err = selector.DoExecute(storage)
		require.NoError(t, err)
		require.Equal(t, "rooms", table.TableName)
		require.Equal(t, []string{"floor", "area"}, table.Columns)
		require.Equal(t, [][]string{{"15", "10"}, {"15", "20"}, {"12", "10"}}, table.Values)
	}

	t.Log("Creator test")
	{
		table.TABLES_PATH = exPath

		storage := table.NewStorage()
		creator := NewCreator("users", []string{"username", "password"})
		require.Equal(t, 0, len(storage.GetTables()))

		table, err := creator.DoExecute(storage)
		require.NoError(t, err)
		require.Equal(t, table.Columns, []string{"username", "password"})
		require.Equal(t, 1, len(storage.GetTables()))

		file, err := os.OpenFile(path.Join(exPath, "users.csv"), os.O_CREATE|os.O_RDONLY, 0600)
		require.NoError(t, err)
		defer file.Close()

		r := csv.NewReader(file)
		r.Comma = ','
		records, err := r.ReadAll()
		require.NoError(t, err)

		require.Equal(t, [][]string{{"username", "password"}}, records)
	}
}
