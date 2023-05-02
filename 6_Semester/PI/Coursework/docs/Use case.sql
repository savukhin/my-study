CREATE TABLE employees(name, surname); -- Ignored
BEGIN trans1;
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT trans1; -- Error: Table not exist

BEGIN trans2;
    CREATE TABLE employees(name, surname);
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT trans2; -- OK


CREATE TABLE employees(name, surname);
COMMIT; -- OK
BEGIN trans3;
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT trans3; -- OK

CREATE TABLE employees(name, surname);
BEGIN trans3;
    INSERT INTO employees VALUEs ("E1", "S1");
    INSERT INTO employees VALUEs ("E2", "S2");
    INSERT INTO employees VALUEs ("E3", "S3");
COMMIT; -- Error (commit trans3): Table not exists
COMMIT trans3; -- -


CREATE TABLE employees(name, surname); -- Ignored
INSERT INTO employees VALUEs ("E1", "S1"); -- Ignored
INSERT INTO employees2 VALUEs ("E2", "S2"); -- Ignored
INSERT INTO employees VALUEs ("E3", "S3"); -- Executed
COMMIT; -- Error: employees not exists 

