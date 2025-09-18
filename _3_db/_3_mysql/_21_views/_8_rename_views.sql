-- In MySQL, views and tables share the same namespace

RENAME TABLE original_view_name 
TO new_view_name;

-- 1. Renaming a view using the RENAME TABLE 
CREATE VIEW productLineSales AS
SELECT 
    productLine, 
    SUM(quantityOrdered) totalQtyOrdered
FROM
    productLines
        INNER JOIN
    products USING (productLine)
        INNER JOIN
    orderdetails USING (productCode)
GROUP BY productLine;

RENAME TABLE productLineSales 
TO productLineQtySales;

SHOW FULL TABLES WHERE table_type = 'VIEW';

-- 2. Renaming a view using the DROP VIEW and CREATE VIEW sequence 
DROP VIEW productLineQtySales;

CREATE VIEW categorySales AS
SELECT 
    productLine, 
    SUM(quantityOrdered) totalQtyOrdered
FROM
    productLines
        INNER JOIN
    products USING (productLine)
        INNER JOIN
    orderDetails USING (productCode)
GROUP BY productLine;