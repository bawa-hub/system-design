-- HAVING clause is used in the SELECT statement to specify filter conditions for a group of rows or aggregates.
-- HAVING clause is often used with the GROUP BY clause to filter groups based on a specified condition. 
-- If you omit the GROUP BY clause, the HAVING clause behaves like the WHERE clause

SELECT 
    select_list
FROM 
    table_name
WHERE 
    search_condition
GROUP BY 
    group_by_expression
HAVING 
    group_condition;

-- HAVING clause evaluates each group returned by the GROUP BY clause
-- if the result is true, the row is included in the result set.

-- Notice that the HAVING clause applies a filter condition to each group of rows, while the WHERE clause applies the filter condition to each individual row.
-- MySQL evaluates the HAVING clause after the FROM, WHERE, SELECT and GROUP BY clauses and before ORDER BY, and LIMIT clauses:


-- Get order numbers, the number of items sold per order, and total sales for each from the orderdetails table
SELECT 
    ordernumber,
    SUM(quantityOrdered) AS itemsCount,
    SUM(priceeach*quantityOrdered) AS total
FROM
    orderdetails
GROUP BY ordernumber;

-- find which order has total sales greater than 1000 
SELECT 
    ordernumber,
    SUM(quantityOrdered) AS itemsCount,
    SUM(priceeach*quantityOrdered) AS total
FROM
    orderdetails
GROUP BY 
   ordernumber
HAVING 
   total > 1000;

-- find orders that have total amounts greater than 1000 and contain more than 600 items
SELECT 
    ordernumber,
    SUM(quantityOrdered) AS itemsCount,
    SUM(priceeach*quantityOrdered) AS total
FROM
    orderdetails
GROUP BY ordernumber
HAVING 
    total > 1000 AND 
    itemsCount > 600;

-- find all orders that already shipped and have a total amount greater than 1500
SELECT 
    a.ordernumber, 
    status, 
    SUM(priceeach*quantityOrdered) total
FROM
    orderdetails a
INNER JOIN orders b 
    ON b.ordernumber = a.ordernumber
GROUP BY  
    ordernumber, 
    status
HAVING 
    status = 'Shipped' AND 
    total > 1500;




-- HAVING clause is only useful when you use it with the GROUP BY clause to generate the output of the high-level reports.
-- For example, you can use the HAVING clause to answer the questions like finding the number of orders this month, this quarter, or this year that have a total amount greater than 10K.