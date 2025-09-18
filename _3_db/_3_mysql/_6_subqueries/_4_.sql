-- MySQL subquery in the FROM clause

-- finds the maximum, minimum, and average number of items in sale orders:

SELECT 
    MAX(items), 
    MIN(items), 
    FLOOR(AVG(items))
FROM
    (SELECT 
        orderNumber, COUNT(orderNumber) AS items
    FROM
        orderdetails
    GROUP BY orderNumber) AS lineitems;

--     +------------+------------+-------------------+
-- | MAX(items) | MIN(items) | FLOOR(AVG(items)) |
-- +------------+------------+-------------------+
-- |         18 |          1 |                 9 |
-- +------------+------------+-------------------+
-- 1 row in set (0.001 sec)