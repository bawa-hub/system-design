-- To delete data from a table
DELETE FROM table_name
WHERE condition;
-- If you omit the WHERE clause, the DELETE statement will delete all rows in the table.
-- DELETE statement returns the number of deleted rows.

-- To delete all rows in a table without the need of knowing how many rows deleted, 
-- you should use the TRUNCATE TABLE statement to get better performance.

-- delete employees whose the officeNumber is 4
DELETE FROM employees 
WHERE officeCode = 4;

--  delete all rows from the employees table
DELETE FROM employees;

-- MySQL DELETE and LIMIT clause

-- sorts customers by customer names alphabetically and deletes the first 10 customers
DELETE FROM customers
ORDER BY customerName
LIMIT 10;


-- selects customers in France, sorts them by credit limit in from low to high, and deletes the first 5 customers
DELETE FROM customers
WHERE country = 'France'
ORDER BY creditLimit
LIMIT 5;