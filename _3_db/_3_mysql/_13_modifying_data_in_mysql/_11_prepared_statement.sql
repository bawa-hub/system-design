-- to make your queries execute faster and more secure.

-- Prior MySQL version 4.1, a query is sent to the MySQL server in the textual format. 
-- In turn, MySQL returns the data to the client using textual protocol
-- MySQL has to fully parse the query and transforms the result set into a string before returning it to the client

-- The prepared statement takes advantage of client/server binary protocol. 
-- It passes the query that contains placeholders (?) to the MySQL Server as:
SELECT * 
FROM products 
WHERE productCode = ?;

-- In order to use MySQL prepared statement, you use three following statements:
--     PREPARE – prepare a statement for execution.
--     EXECUTE – execute a prepared statement prepared by the PREPARE statement.
--     DEALLOCATE PREPARE – release a prepared statement.

-- prepare a statement that returns the product code and name of a product specified by product code:
PREPARE stmt1 FROM 
	'SELECT 
   	    productCode, 
            productName 
	FROM products
        WHERE productCode = ?';

SET @pc = 'S10_1678'; 
EXECUTE stmt1 USING @pc;

SET @pc = 'S12_1099';
EXECUTE stmt1 USING @pc;

DEALLOCATE PREPARE stmt1;