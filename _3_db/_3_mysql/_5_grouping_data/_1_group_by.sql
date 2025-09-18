-- Use the GROUP BY clause to group rows into subgroups.

SELECT 
    c1, c2,..., cn, aggregate_function(ci)
FROM
    table
WHERE
    where_conditions
GROUP BY c1 , c2,...,cn;

-- It works like the DISTINCT operator 

-- GROUP BY clause is often used with an aggregate function to perform calculations and 
-- return a single value for each subgroup.

-- GROUP BY clause vs. DISTINCT clause


-- Examples

-- A. Simple example
SELECT 
    status
FROM
    orders
GROUP BY status;

-- B. With aggregate function

--  To know the number of orders in each status
SELECT 
    status, COUNT(*)
FROM
    orders
GROUP BY status;

-- To get the total amount of all orders by status
SELECT 
    status, 
    SUM(quantityOrdered * priceEach) AS amount
FROM
    orders
INNER JOIN orderdetails 
    USING (orderNumber)
GROUP BY 
    status;

-- Returns the order numbers and the total amount of each order
SELECT 
    orderNumber,
    SUM(quantityOrdered * priceEach) AS total
FROM
    orderdetails
GROUP BY 
    orderNumber;

-- C. With expression

-- gets the total sales for each year.
SELECT 
    YEAR(orderDate) AS year,
    SUM(quantityOrdered * priceEach) AS total
FROM
    orders
INNER JOIN orderdetails 
    USING (orderNumber)
WHERE
    status = 'Shipped'
GROUP BY 
    YEAR(orderDate);
    -- Note that the expression which appears in the SELECT clause must be the same as the one in the GROUP BY clause.
    

-- D. With having
-- To filter the groups returned by GROUP BY clause, you use a  HAVING clause

-- select the total sales of the years after 2003.
SELECT 
    YEAR(orderDate) AS year,
    SUM(quantityOrdered * priceEach) AS total
FROM
    orders
INNER JOIN orderdetails 
    USING (orderNumber)
WHERE
    status = 'Shipped'
GROUP BY 
    year
HAVING 
    year > 2003;


-- The GROUP BY clause vs. DISTINCT clause    
-- If you use the GROUP BY clause in the SELECT statement without using aggregate functions, the GROUP BY clause behaves like the DISTINCT clause.
-- the DISTINCT clause is a special case of the GROUP BY clause.
--  difference between DISTINCT clause and GROUP BY clause is that the GROUP BY clause sorts the result set, whereas the DISTINCT clause does not.

SELECT 
    state
FROM
    customers
GROUP BY state;

-- same as

SELECT DISTINCT
    state
FROM
    customers
ORDER BY 
    state;