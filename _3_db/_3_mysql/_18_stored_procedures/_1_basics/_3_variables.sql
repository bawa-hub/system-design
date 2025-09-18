-- syntax
-- DECLARE variable_name datatype(size) [DEFAULT default_value];

DECLARE totalSale DEC(10,2) DEFAULT 0.0;
DECLARE x, y INT DEFAULT 0;

-- assigning variables
-- SET variable_name = value;

DECLARE total INT DEFAULT 0;
SET total = 10;

-- assign the result of a query to a variable
DECLARE productCount INT DEFAULT 0;

SELECT COUNT(*) 
INTO productCount
FROM products;


-- Variable scopes
-- A variable whose name begins with the @ sign is a session variable. It is available and accessible until the session ends.

DELIMITER $$

CREATE PROCEDURE GetTotalOrder()
BEGIN
	DECLARE totalOrder INT DEFAULT 0;
    
    SELECT COUNT(*) 
    INTO totalOrder
    FROM orders;
    
    SELECT totalOrder;
END$$

DELIMITER ;

CALL GetTotalOrder();