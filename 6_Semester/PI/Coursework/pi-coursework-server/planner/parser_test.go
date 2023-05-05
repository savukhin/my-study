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
		arr := plan.Plan
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
		arr := plan.Plan
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

	t.Log("Insert&Update suite")
	{
		plan, err := ParseOneString("insert into employee (col1, col2, col3) values (val1, val2, val3)")
		arr := plan.Plan
		require.NoError(t, err)
		require.Equal(t, len(arr), 2)

		plan, err = ParseOneString("update employee set username = 'Ivanov' where room == '2'")
		arr = plan.Plan
		require.NoError(t, err)
		require.Equal(t, len(arr), 3)
	}

	t.Log("Test whole query")
	{
		plan, err := ParseFullQuery(`
			BEGIN my_super_transaction;
			UPDATE employee SET room = '15' WHERE surname == 'Ivanov';
			UPDATE employee SET room = '14' WHERE surname == 'Petrov';
			COMMIT my_super_transaction;
		`)
		require.NoError(t, err)

		arr := plan.Plan
		require.Equal(t, len(arr), 8)

		transaction, ok := arr[0].(*processors.BeginTransaction)
		require.True(t, ok)
		require.Equal(t, transaction.Name, "my_super_transaction")

		tabler, ok := arr[1].(*processors.TableGetter)
		require.True(t, ok)
		require.Equal(t, tabler.Table, "employee")
		aggregator, ok := arr[2].(*processors.Aggregator)
		require.True(t, ok)
		require.Equal(t, aggregator.Column, "surname")
		require.Equal(t, aggregator.Sign, "==")
		require.Equal(t, aggregator.Value, "Ivanov")
		updater, ok := arr[3].(*processors.Updater)
		require.True(t, ok)
		require.Equal(t, updater.Column, "room")
		require.Equal(t, updater.NewValue, "15")

		tabler, ok = arr[4].(*processors.TableGetter)
		require.True(t, ok)
		require.Equal(t, tabler.Table, "employee")
		aggregator, ok = arr[5].(*processors.Aggregator)
		require.True(t, ok)
		require.Equal(t, aggregator.Column, "surname")
		require.Equal(t, aggregator.Sign, "==")
		require.Equal(t, aggregator.Value, "Petrov")
		updater, ok = arr[6].(*processors.Updater)
		require.True(t, ok)
		require.Equal(t, updater.Column, "room")
		require.Equal(t, updater.NewValue, "14")

		commit, ok := arr[7].(*processors.CommitTransaction)
		require.True(t, ok)
		require.Equal(t, commit.Name, "my_super_transaction")
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
		table, columns, condition, limit, err := CheckSelector("select surname,   name, room  from employee   ")
		require.NoError(t, err)
		require.Equal(t, table, "employee")
		require.Equal(t, columns, []string{"surname", "name", "room"})
		require.Equal(t, condition.HasWhere, false)
		require.Equal(t, limit.HasLimit, false)
		require.Equal(t, limit.LimitStr, "")

		table, columns, condition, limit, err = CheckSelector("select surname, name, room from employee WHerE 	 room == '4'    LIMIT	 10 ")
		require.NoError(t, err)
		require.Equal(t, table, "employee")
		require.Equal(t, columns, []string{"surname", "name", "room"})
		require.Equal(t, condition.HasWhere, true)
		require.Equal(t, condition.Column, "room")
		require.Equal(t, condition.Sign, "==")
		require.EqualValues(t, condition.ValueStr, "4")
		require.Equal(t, limit.LimitStr, "LIMIT")
		require.Equal(t, limit.HasLimit, true)
		require.EqualValues(t, limit.Limit, 10)

		_, _, _, _, err = CheckSelector("select surname, name, room from employee WHerE 	 room = '4'    LIMIT	 10 ")
		require.Error(t, err)

		table, columns, condition, limit, err = CheckSelector("select * from employee   ")
		require.NoError(t, err)
		require.Equal(t, table, "employee")
		require.Equal(t, columns, []string{"*"})
		require.Equal(t, condition.HasWhere, false)
		require.Equal(t, limit.HasLimit, false)
		require.Equal(t, limit.LimitStr, "")
	}

	t.Log("drop table")
	{
		table, err := CheckDropTable("drop table adf")
		require.NoError(t, err)
		require.Equal(t, table, "adf")

		_, err = CheckDropTable("drop table ")
		require.Error(t, err)
	}

	t.Log("delete rows")
	{
		table, where, err := CheckDeleteRows("delete from adf WhERe room != '4'")
		require.NoError(t, err)
		require.Equal(t, table, "adf")
		require.Equal(t, where.Column, "room")
		require.Equal(t, where.Sign, "!=")
		require.Equal(t, where.ValueStr, "4")

		_, _, err = CheckDeleteRows("delete from WhERe room != '4'")
		require.Error(t, err)

		_, _, err = CheckDeleteRows("delete from adf room != '4'")
		require.Error(t, err)

		_, _, err = CheckDeleteRows("delete from adf WHERE room = '4'")
		require.Error(t, err)
	}

	t.Log("transactions")
	{
		transaction, err := CheckBeginTransaction("begin hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = CheckBeginTransaction("begin")
		require.Error(t, err)

		_, err = CheckBeginTransaction("bEgiN    hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = CheckCommitTransaction("CommIT hiring")
		require.NoError(t, err)
		require.Equal(t, transaction, "hiring")

		_, err = CheckCommitTransaction("Commit hiring limit 10")
		require.Error(t, err)

		err = CheckCommit("CommIT")
		require.NoError(t, err)

		err = CheckCommit("Commit hiring")
		require.Error(t, err)

		err = CheckRollback("RoLLbAck")
		require.NoError(t, err)

		err = CheckCommit("ROLback hiring")
		require.Error(t, err)
	}

	t.Log("updates")
	{
		tableName, setColumnName, setValue, where, err := CheckUpdate("update employee set room = '14' where index == 1")
		require.NoError(t, err)
		require.Equal(t, tableName, "employee")
		require.Equal(t, setColumnName, "room")
		require.Equal(t, setValue, "14")
		require.Equal(t, where.Column, "index")
		require.Equal(t, where.Sign, "==")
		require.EqualValues(t, where.ValueInt, 1)
		require.EqualValues(t, where.Value, "1")
	}

	t.Log("add user")
	{
		username, password, err := CheckAddUser("adD   User USErNAME     	passwOrd    PASsWOrd")
		require.NoError(t, err)
		require.Equal(t, username, "USErNAME")
		require.Equal(t, password, "PASsWOrd")

		_, _, err = CheckAddUser("adD   User USERNAME     	passOrd    PASSWORD")
		require.Error(t, err)
	}

	t.Log("insert")
	{
		tableName, columns, err := CheckInsert("  inSerT   INto   mploy3e2e( col1, col2,COL3,  	COl4  )   values  (  val1,   val2,val3  , val4  )   ")
		require.NoError(t, err)
		require.Equal(t, tableName, "mploy3e2e")
		require.Equal(t, columns, map[string]string{"col1": "val1", "col2": "val2", "COL3": "val3", "COl4": "val4"})

		tableName, columns, err = CheckInsert("inSerT INto mploy3e2e 	( col1 ) values (val1)   ")
		require.NoError(t, err)
		require.Equal(t, tableName, "mploy3e2e")
		require.Equal(t, columns, map[string]string{"col1": "val1"})

		_, _, err = CheckInsert("insert into employee(col1) values (val1,)   ")
		require.Error(t, err)

		_, _, err = CheckInsert("insert into employee (col2, col3) values (val1, val2,)   ")
		require.Error(t, err)

		_, _, err = CheckInsert("insert into employee() values ()   ")
		require.Error(t, err)
	}
}

// func TestWholeQuery(t *testing.T) {
// 	t.Log("Non-transaction parsing test")
// 	{

// 	}
// }
