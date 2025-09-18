-- MySQL subquery with EXISTS and NOT EXISTS

-- finds sales orders whose total values are greater than 60K.

SELECT 
    orderNumber, 
    SUM(priceEach * quantityOrdered) total
FROM
    orderdetails
        INNER JOIN
    orders USING (orderNumber)
GROUP BY orderNumber
HAVING SUM(priceEach * quantityOrdered) > 60000;

-- +-------------+----------+
-- | orderNumber | total    |
-- +-------------+----------+
-- |       10165 | 67392.85 |
-- |       10287 | 61402.00 |
-- |       10310 | 61234.67 |
-- +-------------+----------+
-- 3 rows in set (0.003 sec)


-- find customers who placed at least one sales order with the total value greater than 60K by using the EXISTS operator:

SELECT 
    customerNumber, 
    customerName
FROM
    customers
WHERE
    EXISTS( SELECT 
            orderNumber, SUM(priceEach * quantityOrdered)
        FROM
            orderdetails
                INNER JOIN
            orders USING (orderNumber)
        WHERE
            customerNumber = customers.customerNumber
        GROUP BY orderNumber
        HAVING SUM(priceEach * quantityOrdered) > 60000);

-- +----------------+--------------------------+
-- | customerNumber | customerName             |
-- +----------------+--------------------------+
-- |            148 | Dragon Souveniers, Ltd.  |
-- |            259 | Toms Spezialitäten, Ltd  |
-- |            298 | Vida Sport, Ltd          |
-- +----------------+--------------------------+
-- 3 rows in set (0.002 sec)        