-- MySQL does not have any statement that allows you to directly modify the parameters and body of the stored procedure.

-- you must drop and re-create the stored procedure using the DROP PROCEDURE and CREATE PROCEDURE statements.

DELIMITER $$

CREATE PROCEDURE GetOrderAmount()
BEGIN
    SELECT 
        SUM(quantityOrdered * priceEach) 
    FROM orderDetails;
END$$

DELIMITER ;
