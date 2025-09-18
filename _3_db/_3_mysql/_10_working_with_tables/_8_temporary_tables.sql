-- A temporary table is handy when it is impossible or expensive to query data that requires a single SELECT statement
-- MySQL removes the temporary table automatically when the session ends or the connection is terminated
-- use the  DROP TABLE statement to remove a temporary table explicitly when you are no longer using it.
--  temporary table is only available and accessible to the client that creates it

-- CREATE TEMPORARY TABLE table_name(
--    column1 datatype constraints,
--    column1 datatype constraints,
--    ...,
--    table_constraints
-- );


-- To create a temporary table whose structure is based on an existing table
CREATE TEMPORARY TABLE temp_table_name
SELECT * FROM original_table
LIMIT 0;

-- 1) Creating a temporary table example
CREATE TEMPORARY TABLE credits(
  customerNumber INT PRIMARY KEY, 
  creditLimit DEC(10, 2)
);

INSERT INTO credits(customerNumber, creditLimit)
SELECT 
  customerNumber, 
  creditLimit 
FROM 
  customers 
WHERE 
  creditLimit > 0;

  
-- 2) Creating a temporary table whose structure is based on a query example
CREATE TEMPORARY TABLE top_customers
SELECT p.customerNumber, 
       c.customerName, 
       ROUND(SUM(p.amount),2) sales
FROM payments p
INNER JOIN customers c ON c.customerNumber = p.customerNumber
GROUP BY p.customerNumber
ORDER BY sales DESC
LIMIT 10;

SELECT 
    customerNumber, 
    customerName, 
    sales
FROM
    top_customers
ORDER BY sales;



-- Dropping a temporary table
DROP TEMPORARY TABLE table_name;


-- Checking if a temporary table exists
-- you can create a stored procedure that checks if a temporary table exists or not 
DELIMITER //
CREATE PROCEDURE check_table_exists(table_name VARCHAR(100)) 
BEGIN
    DECLARE CONTINUE HANDLER FOR SQLSTATE '42S02' SET @err = 1;
    SET @err = 0;
    SET @table_name = table_name;
    SET @sql_query = CONCAT('SELECT 1 FROM ',@table_name);
    PREPARE stmt1 FROM @sql_query;
    IF (@err = 1) THEN
        SET @table_exists = 0;
    ELSE
        SET @table_exists = 1;
        DEALLOCATE PREPARE stmt1;
    END IF;
END //
DELIMITER ;

CALL check_table_exists('credits');
SELECT @table_exists;