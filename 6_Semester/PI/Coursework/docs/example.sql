CREATE TABLE employees(name, surname); -- Ignored
INSERT INTO employees VALUEs ("E1", "S1"); -- Ignored
INSERT INTO employees VALUEs ("E2", "S2"); -- Ignored
INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- Added to transaction log till error
ROLLBACK;

Initial commit (empty DB) -> Commit 1 .
Rollback => Revert to initial commit (empty DB)

------------------------------------------------------------------------------

CREATE TABLE employees(name, surname);
COMMIT; -- OK
BEGIN trans1;
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT trans1; -- OK
BEGIN trans2;
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees2 VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT trans2; -- OK
ROLLBACK;

Initial commit (empty DB) => Commit 1 => Commit 2 (trans1) => (trying) commit 3 (trans2) - [error] -> Commit trans1; -
- [rollback] -> Commit 2 


------------------------------------------------------------------------------

CREATE TABLE employees(name, surname);
INSERT INTO employees VALUEs ("E1", "S1");
INSERT INTO employees2 VALUEs ("E2", "S2");
INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- Added to transaction log till error
ROLLBACK;