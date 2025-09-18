-- MySQL supports two kinds of aliases: column alias and table alias.

-- syntax of column alias

-- SELECT 
--    [column_1 | expression] AS descriptive_name
-- FROM table_name;

-- SELECT 
--    [column_1 | expression] AS `descriptive name`
-- FROM 
--    table_name;

SELECT 
    CONCAT_WS(', ', lastName, firstname)
FROM
    employees;


SELECT
   CONCAT_WS(', ', lastName, firstname) AS `Full name`
FROM
   employees;


SELECT
	CONCAT_WS(', ', lastName, firstname) `Full name`
FROM
	employees
ORDER BY
	`Full name`;

SELECT
	orderNumber `Order no.`,
	SUM(priceEach * quantityOrdered) total
FROM
	orderDetails
GROUP BY
	`Order no.`
HAVING
	total > 60000;


-- MySQL alias for tables    
-- table_name AS table_alias

SELECT * FROM employees e;

-- Once a table is assigned an alias, you can refer to the table columns
SELECT 
    e.firstName, 
    e.lastName
FROM
    employees e
ORDER BY e.firstName;


SELECT
	customerName,
	COUNT(o.orderNumber) total
FROM
	customers c
INNER JOIN orders o ON c.customerNumber = o.customerNumber
GROUP BY
	customerName
ORDER BY
	total DESC;


SELECT
	customers.customerName,
	COUNT(orders.orderNumber) total
FROM
	customers
INNER JOIN orders ON customers.customerNumber = orders.customerNumber
GROUP BY
	customerName
ORDER BY
	total DESC    