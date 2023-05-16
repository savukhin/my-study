BEGIN;
    CREATE TABLE musicians(name, surname, band);
    INSERT INTO musicians(surname, band, name) values (Shinoda, LinkinPark, Mike);
    INSERT INTO musicians(name, surname, band) values (Chester, Bennington, LinkinPark);
    INSERT INTO musicians(name, surname, band) values (Maybe, Baby, MaybeBaby);
COMMIT;
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

ROLLBACK;
SELECT * FROM musicians;
