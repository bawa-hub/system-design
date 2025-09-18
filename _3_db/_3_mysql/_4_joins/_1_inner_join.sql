-- select productCode and productName from the products table and 
-- textDescription of product lines from the productlines table

-- SELECT 
--     productCode, 
--     productName, 
--     textDescription
-- FROM
--     products t1
-- INNER JOIN productlines t2 
--     ON t1.productline = t2.productline;


-- shortcut
SELECT 
    productCode, 
    productName, 
    textDescription
FROM
    products
INNER JOIN productlines USING (productline);

-- returns order number, order status, and total sales from the orders and orderdetails tables using the INNER JOIN clause with the GROUP BYclause:

-- SELECT 
--     t1.orderNumber,
--     t1.status,
--     SUM(quantityOrdered * priceEach) total
-- FROM
--     orders t1
-- INNER JOIN orderdetails t2 
--     ON t1.orderNumber = t2.orderNumber
-- GROUP BY orderNumber;

-- or
SELECT 
    orderNumber,
    status,
    SUM(quantityOrdered * priceEach) total
FROM
    orders
INNER JOIN orderdetails USING (orderNumber)
GROUP BY orderNumber;

-- uses two INNER JOIN clauses to join three tables: orders, orderdetails, and products:

SELECT 
    orderNumber,
    orderDate,
    orderLineNumber,
    productName,
    quantityOrdered,
    priceEach
FROM
    orders
INNER JOIN
    orderdetails USING (orderNumber)
INNER JOIN
    products USING (productCode)
ORDER BY 
    orderNumber, 
    orderLineNumber;


-- join four tables orders, orderdetails, customers and products tables

SELECT 
    orderNumber,
    orderDate,
    customerName,
    orderLineNumber,
    productName,
    quantityOrdered,
    priceEach
FROM
    orders
INNER JOIN orderdetails 
    USING (orderNumber)
INNER JOIN products 
    USING (productCode)
INNER JOIN customers 
    USING (customerNumber)
ORDER BY 
    orderNumber, 
    orderLineNumber;

-- find the sales price of the product whose code is S10_1678 that is less than the manufacturer’s suggested retail price (MSRP) for that product

SELECT 
    orderNumber, 
    productName, 
    msrp, 
    priceEach
FROM
    products p
INNER JOIN orderdetails o 
   ON p.productcode = o.productcode
      AND p.msrp > o.priceEach
WHERE
    p.productcode = 'S10_1678';    