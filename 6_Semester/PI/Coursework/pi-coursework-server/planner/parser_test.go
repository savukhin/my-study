package planner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	t.Log("Create table suite")
	{
		plan, err := Parse("create table employees (surname, name, room)")
		arr := plan.plan
		require.NoError(t, err)
		require.Equal(t, len(arr), 1)
		require.Equal(t, arr[0].GetName(), TableCreatorName)

		creator := arr[0].(*TableCreator)
		fmt.Println("CREATOR = ", creator)
		fmt.Println("CREATOR = ", creator.columns)
		fmt.Println("CREATOR = ", creator.Table)
		require.Equal(t, []string{"surname", "name", "room"}, creator.columns)
		require.Equal(t, "employees", creator.Table)
	}
}

func TestParserFail(t *testing.T) {
	t.Log("Create table suite")
	{
		_, err := Parse("create table emlpoyees (surname name room)")
		require.Error(t, err)

		_, err = Parse("create table emlpoyees (surname, name,)")
		require.Error(t, err)

		_, err = Parse("createtable emlpoyees (surname, name, room)")
		require.Error(t, err)

		_, err = Parse("create tabl emlpoyees (surname, name, room)")
		require.Error(t, err)

		_, err = Parse("create table (surname, name, room)")
		require.Error(t, err)

		_, err = Parse("create tabl emlpoyees")
		require.Error(t, err)

		_, err = Parse("create tabl emlpoyees ()")
		require.Error(t, err)

		_, err = Parse("create tabl emlpoyees (,)")
		require.Error(t, err)
	}
}
