-- left join two tables
-- use the LEFT JOIN clause to find all customers and their orders:

-- SELECT 
--     customers.customerNumber, 
--     customerName, 
--     orderNumber, 
--     status
-- FROM
--     customers
-- LEFT JOIN orders ON 
--     orders.customerNumber = customers.customerNumber;

-- or
SELECT
    c.customerNumber,
    customerName,
    orderNumber,
    status
FROM
    customers c
LEFT JOIN orders o 
    ON c.customerNumber = o.customerNumber;  

-- or
-- SELECT
-- 	customerNumber,
-- 	customerName,
-- 	orderNumber,
-- 	status
-- FROM
-- 	customers
-- LEFT JOIN orders USING (customerNumber);
    
-- customers is the left table and orders is the right table
-- LEFT JOIN clause returns all customers including the customers who have no order. 
-- If a customer has no order, the values in the column orderNumber and status are NULL


-- If you replace the LEFT JOIN clause by the INNER JOIN clause, you will get the only customers who have at least one order


-- left join to find unmatched rows
-- use the LEFT JOIN to find customers who have no order:

SELECT 
    c.customerNumber, 
    c.customerName, 
    o.orderNumber, 
    o.status
FROM
    customers c
LEFT JOIN orders o 
    ON c.customerNumber = o.customerNumber
WHERE
    orderNumber IS NULL;



-- left join three tables
-- use two LEFT JOIN clauses to join the three tables: employees, customers, and payments

SELECT 
    lastName, 
    firstName, 
    customerName, 
    checkNumber, 
    amount
FROM
    employees
LEFT JOIN customers ON 
    employeeNumber = salesRepEmployeeNumber
LEFT JOIN payments ON 
    payments.customerNumber = customers.customerNumber
ORDER BY 
    customerName, 
    checkNumber;

-- return the order and its line items of the order number 10123.

SELECT 
    o.orderNumber, 
    customerNumber, 
    productCode
FROM
    orders o
LEFT JOIN orderDetails 
    USING (orderNumber)
WHERE
    orderNumber = 10123;

-- returns all orders but only the order 10123 will have line items associated with it 

SELECT 
    o.orderNumber, 
    customerNumber, 
    productCode
FROM
    orders o
LEFT JOIN orderDetails d 
    ON o.orderNumber = d.orderNumber AND 
       o.orderNumber = 10123;