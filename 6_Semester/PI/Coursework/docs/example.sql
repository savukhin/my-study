CREATE TABLE employees(name, surname); -- Ignored
INSERT INTO employees VALUEs ("E1", "S1"); -- Ignored
INSERT INTO employees VALUEs ("E2", "S2"); -- Ignored
INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- Added to transaction log till error
ROLLBACK;

Initial commit (empty DB) -> Commit 1 .
Rollback => Revert to initial commit (empty DB)

------------------------------------------------------------------------------

BEGIN; --- t1
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- OK
BEGIN ; --- t2
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees2 VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- OK
ROLLBACK; -- OK

Initial commit (empty DB) => Commit 1 => Commit 2 (trans1) => (trying) commit 3 (trans2) - [error] -> Commit trans1; -
- [rollback] -> Commit 2 

--------

[transaction][operations]IExecutor executors

for transaction in transactions
    trans = executors[transaction]
    tmp_storage = current_storage.Copy()

    if len(trans) == 1 and trans[0] is ROLLBACK
        -- Get array of changes before last COMMIT;
        -- [0] is the oldest, [-1] is the newest
        last_changes = transaction_log.get_rollback() 

        for change in last_changes[::-1] -- Revert order
            tmp_storage = change.revert(tmp_storage)
        if success
            WriteToTransactionLog(trans)
            current_storage = tmp_storage
        else 
            fatal("Data corruption")

    for operation in trans
        tmp_storage = operation.DoExecute(tmp_storage)
    if success
        WriteToTransactionLog(trans)
        current_storage = tmp_storage
    else
        throw Error("Cannot do transaction", transaction)


------------------------------------------------------------------------------

transaction_log:
    Stack<ComplexTransaction> activeTransactions
    []ComplexTransaction complexTransactions

AbstractTransaction
- ComplexTransaction
    - Updates current_storage and writes to transaction_log (physical and cache)

function Eval(storage, *transaction_log, *current_storage)
    for operation in this.operations
        tmp_storage, err = operation.DoExecute(tmp_storage)

    if success
        transaction_log.write(this.operations)
        current_storage = tmp_storage
    else
        throw Error("Cannot do transaction", transaction)


- RollbackTransaction
    - Writes to transaction_log (physical and cache) and cached tables

function Eval (storage, *transaction_log, *current_storage)
    transaction = transaction_log.GetLastActiveComplexTransaction() -- Not rollbacked
    transaction_log.AddRollbackEvent(transaction.id)

    transaction_rollbacked = transaction.rollback() -- Get the reverted sequence of complex transaction
    transaction_rollbacked.Eval(storage, transaction_log)


- WriteTransaction
    - Reads to transaction_log (physical and cache) and writes to physical storage

function Eval(storage, *transaction_log, *current_storage)
    start = 0
    storage_loaded := LoadStorage()

    for i = len(transaction_log.Events) - 1; i >= 0; i++ 
        event = transaction_log.Events
        
        if event.type is WRITE_TRANSACTION
            start = i
            break

    for i = start; i < len(transaction_log.Events); i++
        if event.type is ROLLBACK
            rollback = event as ROLLBACK_EVENT
            trans = transaction_log.find_transaction(rollback.rollbacked_transaction_id)

            rollbacked = trans.rollback
            rollbacked.EvalWrite(storage_loaded)
            continue

        
        if events.type IS COMPLEX_TRANSACTION
            trans = event as COMPLEX_TRANSACTION
            rollbacked.EvalWrite(storage_loaded)
        
        if events.type IS WRITE_TRANSACTION
            fatal("Data corrupted")
        



if query is SelectQuery 
    table = Aggregate(query)
    return json.dump(table)

transactions = ParseWholeQuery(query)
err = mainExecutor(transactions)
if err is nil 
    return "OK"
else
    return err


function mainExecutor(transactions)
    for transaction in transactions
        err = transaction.Eval(storage, transaction_log)


------------------------------------------------------------------------------
BEGIN; -- t1
INSERT ...;
UPDATE ...;
DELETE ...;
BEGIN; -- t2
CREATE ...;
INSERT ...;
COMMIT;
ROLLBACK; -- revert t2 |  -- r_t2
WRITE; -- w1
ROLLBACK; -- revert t1 |  -- r_t1
BEGIN; -- t3
INSERT ...;
UPDATE ...;
DELETE ...;
COMMIT;
ROLLBACK; -- revert t3 |  -- r_t3


transaction_id,event_type
t1,     insert
t1,     update
t2,     create
t2,     insert
r_t2,   rollback
w1,     write
r_t1,   rollback
t3,     insert
t3,     udpate
t3,     delete
r_t3,   rollback

How to find what to write
1. We know 

------------------------------------------------------------------------------
BEGIN; -- t1
CREATE ...;
INSERT ...;
COMMIT;
WRITE; -- w1
BEGIN; -- t2
INSERT ...;
UPDATE ...;
DELETE ...;
COMMIT;
ROLLBACK; -- revert t2 |  -- r_t2
ROLLBACK; -- revert t1 |  -- r_t1
ROLLBACK; -- revert t0 |  -- r_t0
                       |
Every ROLLBACK must  <-|
be in transacton log <-|

=>
transaction_id,event_type
t1,     create
t1,     insert
w1,     write
t2,     insert
t2,     udpate
t2,     delete
r_t2,   rollback
r_t1,   rollback


------------------------------------------------------------------------------

CREATE TABLE employees(name, surname);
INSERT INTO employees VALUEs ("E1", "S1");
INSERT INTO employees2 VALUEs ("E2", "S2");
INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- Added to transaction log till error
ROLLBACK;