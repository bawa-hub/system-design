-- syntax
SELECT 
   select_list
FROM 
   table_name
ORDER BY 
   column1 [ASC|DESC], 
   column2 [ASC|DESC];

-- sort the customers by their last names in ascending order
SELECT
	contactLastname,
	contactFirstname
FROM
	customers
ORDER BY
	contactLastname; -- ASC is by default


--  sort the customers by the last name in descending order and then by the first name in ascending order
SELECT 
    contactLastname, 
    contactFirstname
FROM
    customers
ORDER BY 
	contactLastname DESC , 
	contactFirstname ASC;

-- Using MySQL ORDER BY clause to sort data using a custom list
FIELD(str, str1, str2, ...)
-- FIELD() function returns the position of the str in the str1, str2, … list. 
-- If the str is not in the list, the FIELD() function returns 0

SELECT FIELD('A', 'A', 'B','C');
Output:

-- +--------------------------+
-- | FIELD('A', 'A', 'B','C') |
-- +--------------------------+
-- |                        1 |
-- +--------------------------+
-- 1 row in set (0.00 sec)

-- sort the sales orders based on their statuses
SELECT 
    orderNumber, status
FROM
    orders
ORDER BY FIELD(status,
        'In Process',
        'On Hold',
        'Cancelled',
        'Resolved',
        'Disputed',
        'Shipped');

+-------------+------------+
| orderNumber | status     |
+-------------+------------+
|       10425 | In Process |
|       10421 | In Process |
|       10422 | In Process |
|       10420 | In Process |
|       10424 | In Process |
|       10423 | In Process |
|       10414 | On Hold    |
|       10401 | On Hold    |
|       10334 | On Hold    |
|       10407 | On Hold    |
...


-- In MySql, NULL comes before non-NULL values

-- sort employees by values in the reportsTo column in ascending order
SELECT 
    firstName, lastName, reportsTo
FROM
    employees
ORDER BY reportsTo;

-- Output:

-- +-----------+-----------+-----------+
-- | firstName | lastName  | reportsTo |
-- +-----------+-----------+-----------+
-- | Diane     | Murphy    |      NULL |
-- | Mary      | Patterson |      1002 |
-- | Jeff      | Firrelli  |      1002 |
-- | William   | Patterson |      1056 |
-- | Gerard    | Bondur    |      1056 |
-- ...

-- sort employees by values in the reportsTo column in descending order
SELECT 
    firstName, lastName, reportsTo
FROM
    employees
ORDER BY reportsTo DESC;

-- Output:

-- +-----------+-----------+-----------+
-- | firstName | lastName  | reportsTo |
-- +-----------+-----------+-----------+
-- | Yoshimi   | Kato      |      1621 |
-- | Leslie    | Jennings  |      1143 |
-- | Leslie    | Thompson  |      1143 |
-- | Julie     | Firrelli  |      1143 |
-- | ....
-- | Mami      | Nishi     |      1056 |
-- | Mary      | Patterson |      1002 |
-- | Jeff      | Firrelli  |      1002 |
-- | Diane     | Murphy    |      NULL |
-- +-----------+-----------+-----------+
-- 23 rows in set (0.00 sec)