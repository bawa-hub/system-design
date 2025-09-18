-- MySQL derived tables

-- derived table is a virtual table returned from a SELECT statement
-- derived table is similar to a temporary table, but using a derived table in the SELECT statement is much simpler than a temporary table because 
-- it does not require creating the temporary table.
-- When a stand-alone subquery is used in the FROM clause of a SELECT statement, it is also called a derived table.

-- syntax
-- SELECT 
--     select_list
-- FROM
--     (SELECT 
--         select_list
--     FROM
--         table_1) derived_table_name
-- WHERE 
--     derived_table_name.c1 > 0;

-- gets the top five products by sales revenue in 2003 from the orders and orderdetails tables
SELECT 
    productCode, 
    ROUND(SUM(quantityOrdered * priceEach)) sales
FROM
    orderdetails
        INNER JOIN
    orders USING (orderNumber)
WHERE
    YEAR(shippedDate) = 2003
GROUP BY productCode
ORDER BY sales DESC
LIMIT 5;

-- You can use the result of this query as a derived table and join it with the products table as follows
SELECT 
    productName, sales
FROM
    (SELECT 
        productCode, 
        ROUND(SUM(quantityOrdered * priceEach)) sales
    FROM
        orderdetails
    INNER JOIN orders USING (orderNumber)
    WHERE
        YEAR(shippedDate) = 2003
    GROUP BY productCode
    ORDER BY sales DESC
    LIMIT 5) top5products2003
INNER JOIN
    products USING (productCode);



-- classify the customers who bought products in 2003 into 3 groups: platinum, gold, and silver. 
-- And you need to know the number of customers in each group with the following conditions
--    Platinum customers who have orders with the volume greater than 100K.
--    Gold customers who have orders with the volume between 10K and 100K.
--    Silver customers who have orders with the volume less than 10K.

SELECT 
    customerNumber,
    ROUND(SUM(quantityOrdered * priceEach)) sales,
    (CASE
        WHEN SUM(quantityOrdered * priceEach) < 10000 THEN 'Silver'
        WHEN SUM(quantityOrdered * priceEach) BETWEEN 10000 AND 100000 THEN 'Gold'
        WHEN SUM(quantityOrdered * priceEach) > 100000 THEN 'Platinum'
    END) customerGroup
FROM
    orderdetails
        INNER JOIN
    orders USING (orderNumber)
WHERE
    YEAR(shippedDate) = 2003
GROUP BY customerNumber;

--  use above query as the derived table and perform grouping as follows
SELECT 
    customerGroup, 
    COUNT(cg.customerGroup) AS groupCount
FROM
    (SELECT 
        customerNumber,
            ROUND(SUM(quantityOrdered * priceEach)) sales,
            (CASE
                WHEN SUM(quantityOrdered * priceEach) < 10000 THEN 'Silver'
                WHEN SUM(quantityOrdered * priceEach) BETWEEN 10000 AND 100000 THEN 'Gold'
                WHEN SUM(quantityOrdered * priceEach) > 100000 THEN 'Platinum'
            END) customerGroup
    FROM
        orderdetails
    INNER JOIN orders USING (orderNumber)
    WHERE
        YEAR(shippedDate) = 2003
    GROUP BY customerNumber) cg
GROUP BY cg.customerGroup;    
