package planner

import (
	"pi-coursework-server/processors"
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

		creator, ok := arr[0].(*processors.TableCreator)
		require.True(t, ok)
		require.Equal(t, []string{"surname", "name", "room"}, creator.Columns)
		require.Equal(t, "employees", creator.Table)
	}

	t.Log("Selector suite")
	{
		plan, err := Parse("select surname, name, room from employee WHerE 	 room == '4'    LIMIT	 10 ")
		arr := plan.plan
		require.NoError(t, err)
		require.Equal(t, len(arr), 4)

		tableGetter, ok := arr[0].(*processors.TableGetter)
		require.True(t, ok)
		aggregator, ok := arr[1].(*processors.Aggregator)
		require.True(t, ok)
		selector, ok := arr[2].(*processors.Selector)
		require.True(t, ok)
		limiter, ok := arr[3].(*processors.Limiter)
		require.True(t, ok)

		require.Equal(t, []string{"surname", "name", "room"}, selector.Columns)
		require.Equal(t, "employee", tableGetter.Table)
		require.Equal(t, "room", aggregator.Column)
		require.Equal(t, "4", aggregator.Value)
		require.Equal(t, "==", aggregator.Sign)
		require.EqualValues(t, 10, limiter.Count)
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

func TestCheckSelector(t *testing.T) {
	table, columns, condition, limit, err := checkSelector("select surname,   name, room  from employee   ")
	require.NoError(t, err)
	require.Equal(t, table, "employee")
	require.Equal(t, columns, []string{"surname", "name", "room"})
	require.Equal(t, condition.HasWhere, false)
	require.Equal(t, limit.HasLimit, false)
	require.Equal(t, limit.LimitStr, "")

	table, columns, condition, limit, err = checkSelector("select surname, name, room from employee WHerE 	 room == '4'    LIMIT	 10 ")
	require.NoError(t, err)
	require.Equal(t, table, "employee")
	require.Equal(t, columns, []string{"surname", "name", "room"})
	require.Equal(t, condition.HasWhere, true)
	require.Equal(t, condition.Column, "room")
	require.Equal(t, condition.Sign, "==")
	require.Equal(t, condition.Value, "4")
	require.Equal(t, limit.LimitStr, "LIMIT")
	require.Equal(t, limit.HasLimit, true)
	require.EqualValues(t, limit.Limit, 10)

	table, columns, condition, limit, err = checkSelector("select surname, name, room from employee WHerE 	 room = '4'    LIMIT	 10 ")
	require.Error(t, err)
}
