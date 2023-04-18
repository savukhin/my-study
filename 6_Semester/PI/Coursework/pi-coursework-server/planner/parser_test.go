package planner

import (
	"pi-coursework-server/processors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	t.Log("Create table suite")
	{
		plan, err := ParseOneString("create table employees (surname, name, room)")
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
		plan, err := ParseOneString("select surname, name, room from employee WHerE 	 room == '4'    LIMIT	 10 ")
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
		_, err := ParseOneString("create table emlpoyees (surname name room)")
		require.Error(t, err)

		_, err = ParseOneString("create table emlpoyees (surname, name,)")
		require.Error(t, err)

		_, err = ParseOneString("createtable emlpoyees (surname, name, room)")
		require.Error(t, err)

		_, err = ParseOneString("create tabl emlpoyees (surname, name, room)")
		require.Error(t, err)

		_, err = ParseOneString("create table (surname, name, room)")
		require.Error(t, err)

		_, err = ParseOneString("create tabl emlpoyees")
		require.Error(t, err)

		_, err = ParseOneString("create tabl emlpoyees ()")
		require.Error(t, err)

		_, err = ParseOneString("create tabl emlpoyees (,)")
		require.Error(t, err)
	}
}

func TestCheckers(t *testing.T) {
	t.Log("check selector")
	{
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

		_, _, _, _, err = checkSelector("select surname, name, room from employee WHerE 	 room = '4'    LIMIT	 10 ")
		require.Error(t, err)
	}

	t.Log("drop table")
	{
		table, err := checkDropTable("drop table adf")
		require.NoError(t, err)
		require.Equal(t, table, "adf")

		_, err = checkDropTable("drop table ")
		require.Error(t, err)
	}

	t.Log("delete rows")
	{
		table, where, err := checkDeleteRows("delete from adf WhERe room != '4'")
		require.NoError(t, err)
		require.Equal(t, table, "adf")
		require.Equal(t, where.Column, "room")
		require.Equal(t, where.Sign, "!=")
		require.Equal(t, where.Value, "4")

		_, _, err = checkDeleteRows("delete from WhERe room != '4'")
		require.Error(t, err)

		_, _, err = checkDeleteRows("delete from adf room != '4'")
		require.Error(t, err)

		_, _, err = checkDeleteRows("delete from adf WHERE room = '4'")
		require.Error(t, err)
	}

	t.Log("transactions")
	{
		transaction, err := checkBeginTransaction("begin hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = checkBeginTransaction("begin")
		require.Error(t, err)

		_, err = checkBeginTransaction("bEgiN    hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = checkCommitTransaction("CommIT hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = checkCommitTransaction("Commit hiring limit 10")
		require.Error(t, err)

		err = checkCommit("CommIT")
		require.NoError(t, err)

		err = checkCommit("Commit hiring")
		require.Error(t, err)

		err = checkRollback("RoLLbAck")
		require.NoError(t, err)

		err = checkCommit("ROLback hiring")
		require.Error(t, err)
	}

	t.Log("updates")
	{
		// transaction, err := checkUpdate("update employee set room = '14' where index == 1")
	}
}
