-- MySQL EXISTS operator

-- EXISTS operator is a Boolean operator that returns either true or false
-- EXISTS operator is often used to test for the existence of rows returned by the subquery

-- syntax
-- SELECT 
--     select_list
-- FROM
--     a_table
-- WHERE
--     [NOT] EXISTS(subquery);

-- subquery returns at least one row, the EXISTS operator returns true, otherwise, it returns false.
-- EXISTS operator terminates further processing immediately once it finds a matching row, which can help improve the performance of the query.



-- find the customer who has at least one order:
SELECT 
    customerNumber, 
    customerName
FROM
    customers
WHERE
    EXISTS(
	SELECT 
            1
        FROM
            orders
        WHERE
            orders.customernumber = customers.customernumber);

-- find customers who do not have any orders:            
SELECT 
    customerNumber, 
    customerName
FROM
    customers
WHERE
    NOT EXISTS( 
	SELECT 
            1
        FROM
            orders
        WHERE
            orders.customernumber = customers.customernumber
	);


-- finds employees who work at the office in San Franciso
SELECT 
    employeenumber, 
    firstname, 
    lastname, 
    extension
FROM
    employees
WHERE
    EXISTS( 
        SELECT 
            1
        FROM
            offices
        WHERE
            city = 'San Francisco' AND 
           offices.officeCode = employees.officeCode);


-- adds the number 1 to the phone extension of employees who work at the office in San Francisco:           
UPDATE employees 
SET 
    extension = CONCAT(extension, '1')
WHERE
    EXISTS( 
        SELECT 
            1
        FROM
            offices
        WHERE
            city = 'San Francisco'
                AND offices.officeCode = employees.officeCode);


-- MySQL INSERT EXISTS

-- archive customers who don’t have any sales order in a separate table
-- 1. First, create a new table for archiving the customers by copying the structure from the customers table
CREATE TABLE customers_archive 
LIKE customers;

-- 2. Second, insert customers who do not have any sales order into the customers_archive table
INSERT INTO customers_archive
SELECT * 
FROM customers
WHERE NOT EXISTS( 
   SELECT 1
   FROM
       orders
   WHERE
       orders.customernumber = customers.customernumber
);

-- MySQL DELETE EXISTS
-- 3. finally delete the customers that exist in the customers_archive table from the customers table
DELETE FROM customers
WHERE EXISTS( 
    SELECT 
        1
    FROM
        customers_archive a
    
    WHERE
        a.customernumber = customers.customerNumber);


-- MySQL EXISTS operator vs. IN operator

-- find the customer who has placed at least one order, you can use the IN operator
SELECT 
    customerNumber, 
    customerName
FROM
    customers
WHERE
    customerNumber IN (
        SELECT 
            customerNumber
        FROM
            orders);


-- for performance between exists and in operator:
-- https://www.mysqltutorial.org/mysql-exists/            