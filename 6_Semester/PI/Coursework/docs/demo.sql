BEGIN a;
    CREATE TABLE musicians(name, surname, band);
BEGIN b;
    CREATE TABLE musicians(name, surname, band);
SELECT * FROM musicians;

COMMIT a;
ROLLBACK a;


BEGIN;
    CREATE TABLE musicians(name, surname, band);
    INSERT INTO musicians(surname, band, name) values (Shinoda, LinkinPark, Mike);
    INSERT INTO musicians(name, surname, band) values (Chester, Bennington, LinkinPark);
    INSERT INTO musicians(name, surname, band) values (Maybe, Baby, MaybeBaby);
COMMIT;
SELECT * FROM musicians;

BEGIN;
    INSERT INTO musicians(name, surname, band) values (Ozzy, Osbourne, BlackSabbath);
COMMIT;
SELECT * FROM musicians;

ROLLBACk;
SELECT * FROM musicians;

BEGIN;
    DELETE FROM musicians where name == 'Maybe';
COMMIT;
SELECT * FROM musicians;

ROLLBACK;
SELECT * FROM musicians WHERE name == 'Maybe';

BEGIN;
    UPDATE musicians SET band = 'BabyMaybe' WHERE name == 'Maybe';
COMMIT;
SELECT * FROM musicians;

BEGIN;
    UPDATE musicians SET band = 'BabyMaybe' WHERE name == 'Maybe1';
COMMIT;
SELECT * FROM musicians;

ROLLBACK;
SELECT * FROM musicians;

SELECT name, surname FROM musicians;

SELECT name, surname FROM musicians WHERE name != 'Mike';

SELECT name, surname FROM musicians WHERE name != 'Mike' LIMIT 1;

BEGIN;
    DROP TABLE musicians;
COMMIT;
SELECT * FROM musicians;

ROLLBACK;
SELECT * FROM musicians;

ДОЛЖНЫ БЫТЬ ИМЕНОВАННЫЕ ТРАНЗАКЦИИ