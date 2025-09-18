DROP [TEMPORARY] TABLE [IF EXISTS] table_name [, table_name] ...
[RESTRICT | CASCADE]

-- DROP TABLE statement removes a table and its data permanently from the database

-- To execute the DROP TABLE statement, you must have DROP privileges for the table that you want to remove.

-- Using MySQL DROP TABLE to drop a single table

CREATE TABLE insurances (
    id INT AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    effectiveDate DATE NOT NULL,
    duration INT NOT NULL,
    amount DEC(10 , 2 ) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE insurances;

-- Using MySQL DROP TABLE to drop multiple tables

CREATE TABLE CarAccessories (
    id INT AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    price DEC(10 , 2 ) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE CarGadgets (
    id INT AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    price DEC(10 , 2 ) NOT NULL,
    PRIMARY KEY(id)
);

DROP TABLE CarAccessories, CarGadgets;

--  Using MySQL DROP TABLE to drop a non-existing table

DROP TABLE aliens;
-- Error Code: 1051. Unknown table 'classicmodels.aliens'

DROP TABLE IF EXISTS aliens;
-- 0 row(s) affected, 1 warning(s): 1051 Unknown table 'classicmodels.aliens'

SHOW WARNINGS;

-- DROP TABLE based on a pattern


-- MySQL does not have the DROP TABLE LIKE statement that can remove tables based on pattern matching:
DROP TABLE LIKE '%pattern%'

CREATE TABLE test1(
  id INT AUTO_INCREMENT,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS test2 LIKE test1;
CREATE TABLE IF NOT EXISTS test3 LIKE test1;

-- remove all test* tables.

-- set table schema and pattern matching for tables
SET @schema = 'classicmodels';
SET @pattern = 'test%';

-- construct dynamic sql (DROP TABLE tbl1, tbl2...;)
SELECT CONCAT('DROP TABLE ',GROUP_CONCAT(CONCAT(@schema,'.',table_name)),';')
INTO @droplike
FROM information_schema.tables
WHERE @schema = database()
AND table_name LIKE @pattern;

-- display the dynamic sql statement
SELECT @droplike;

-- execute dynamic sql
PREPARE stmt FROM @droplike;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;