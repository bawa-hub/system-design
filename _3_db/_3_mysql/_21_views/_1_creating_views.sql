-- syntax
CREATE [OR REPLACE] VIEW [db_name.] view_name [(column_list)] AS
select - statement;

-- name of view cannot be same as existing table

-- create a view that represents total sales per order.
CREATE VIEW salePerOrder AS
SELECT orderNumber,
    SUM(quantityOrdered * priceEach) total
FROM orderDetails
GROUP by orderNumber
ORDER BY total DESC;

-- show all tables include views
SHOW FULL TABLES;

-- query total sales for each sales order
SELECT * FROM salePerOrder;

-- MySQL allows you to create a view based on another view

-- create a view called bigSalesOrder based on the salesPerOrder view to show every sales order whose total is greater than 60,000 
CREATE VIEW bigSalesOrder AS
    SELECT 
        orderNumber, 
        ROUND(total,2) as total
    FROM
        salePerOrder
    WHERE
        total > 60000;

SELECT 
    orderNumber, 
    total
FROM
    bigSalesOrder;



-- Creating a view with join
CREATE OR REPLACE VIEW customerOrders AS
SELECT 
    orderNumber,
    customerName,
    SUM(quantityOrdered * priceEach) total
FROM
    orderDetails
INNER JOIN orders o USING (orderNumber)
INNER JOIN customers USING (customerNumber)
GROUP BY orderNumber;

-- Creating a view with a subquery

-- products whose buy prices are higher than the average price of all products.
CREATE VIEW aboveAvgProducts AS
    SELECT 
        productCode, 
        productName, 
        buyPrice
    FROM
        products
    WHERE
        buyPrice > (
            SELECT 
                AVG(buyPrice)
            FROM
                products)
    ORDER BY buyPrice DESC;


-- Creating a view with explicit view columns
CREATE VIEW customerOrderStats (
   customerName , 
   orderCount
) 
AS
    SELECT 
        customerName, 
        COUNT(orderNumber)
    FROM
        customers
            INNER JOIN
        orders USING (customerNumber)
    GROUP BY customerName;

SELECT 
    customerName,
    orderCount
FROM
    customerOrderStats
ORDER BY 
	orderCount, 
    customerName;