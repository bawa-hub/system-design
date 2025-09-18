-- MySQL REPLACE statement works as follows:

-- Step 1. Insert a new row into the table, if a duplicate key error occurs.
-- Step 2. If the insertion fails due to a duplicate-key error occurs:
        --    Delete the conflicting row that causes the duplicate key error from the table.
        --    Insert the new row into the table again


-- syntax
REPLACE [INTO] table_name(column_list)
VALUES(value_list);

-- Using MySQL REPLACE to insert a new row

CREATE TABLE cities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    population INT NOT NULL
);

INSERT INTO cities(name,population)
VALUES('New York',8008278),
	  ('Los Angeles',3694825),
	  ('San Diego',1223405);

SELECT * FROM cities;

REPLACE INTO cities(id,population)
VALUES(2,3696820);      

SELECT * FROM cities;

-- Using MySQL REPLACE statement to update a row

-- syntax
REPLACE INTO table
SET column1 = value1,
    column2 = value2;

REPLACE INTO cities
SET id = 4,
    name = 'Phoenix',
    population = 1768980;


-- Using MySQL REPLACE to insert data from a SELECT statement

REPLACE INTO table_1(column_list)
SELECT column_list
FROM table_2
WHERE where_condition;


REPLACE INTO 
    cities(name,population)
SELECT 
    name,
    population 
FROM 
   cities 
WHERE id = 1;