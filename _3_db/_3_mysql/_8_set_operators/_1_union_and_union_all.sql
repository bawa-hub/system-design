-- UNION operator allows you to combine two or more result sets of queries into a single result set

-- syntax
-- SELECT column_list
-- UNION [DISTINCT | ALL]
-- SELECT column_list
-- UNION [DISTINCT | ALL]
-- SELECT column_list
-- ...

-- basic rules to follow to combine two or more queries results
-- a. number and the orders of columns that appear in all SELECT statements must be the same.
-- b. data types of columns must be the same or compatible.

-- UNION operator removes duplicate rows even if you don’t specify the DISTINCT operator explicitly.


-- basic example
DROP TABLE IF EXISTS t1;
DROP TABLE IF EXISTS t2;

CREATE TABLE t1 (
    id INT PRIMARY KEY
);

CREATE TABLE t2 (
    id INT PRIMARY KEY
);

INSERT INTO t1 VALUES (1),(2),(3);
INSERT INTO t2 VALUES (2),(3),(4);

SELECT id
FROM t1
UNION
SELECT id
FROM t2;

-- +----+
-- | id |
-- +----+
-- |  1 |
-- |  2 |
-- |  3 |
-- |  4 |
-- +----+
-- 4 rows in set (0.00 sec)

SELECT id
FROM t1
UNION ALL
SELECT id
FROM t2;

-- +----+
-- | id |
-- +----+
-- |  1 |
-- |  2 |
-- |  3 |
-- |  2 |
-- |  3 |
-- |  4 |
-- +----+
-- 6 rows in set (0.00 sec)

-- union vs join
-- JOIN combines result sets horizontally, a UNION appends result set vertically.

-- combine the first name and last name of employees and customers into a single result set
SELECT 
    firstName, 
    lastName
FROM
    employees 
UNION 
SELECT 
    contactFirstName, 
    contactLastName
FROM
    customers;


-- sort the result set of a union
SELECT 
    concat(firstName,' ',lastName) fullname
FROM
    employees 
UNION SELECT 
    concat(contactFirstName,' ',contactLastName)
FROM
    customers
ORDER BY fullname;


-- differentiate between employees and customers
SELECT 
    CONCAT(firstName, ' ', lastName) fullname, 
    'Employee' as contactType
FROM
    employees 
UNION SELECT 
    CONCAT(contactFirstName, ' ', contactLastName),
    'Customer' as contactType
FROM
    customers
ORDER BY 
    fullname