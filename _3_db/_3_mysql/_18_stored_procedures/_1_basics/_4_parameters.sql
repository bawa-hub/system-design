--  parameter in a stored procedure has one of three modes: IN,OUT, or INOUT.


-- syntax
-- [IN | OUT | INOUT] parameter_name datatype[(length)]

-- IN param

-- IN is the default mode. When you define an IN parameter in a stored procedure, the calling program has to pass an argument to the stored procedure
-- the value of an IN parameter is protected
-- It means that even you change the value of the IN parameter inside the stored procedure, its original value is unchanged after the stored procedure ends.
-- the stored procedure only works on the copy of the IN parameter.


DELIMITER //

CREATE PROCEDURE GetOfficeByCountry(
	IN countryName VARCHAR(255)
)
BEGIN
	SELECT * 
 	FROM offices
	WHERE country = countryName;
END //

DELIMITER ;

CALL GetOfficeByCountry('USA');
CALL GetOfficeByCountry('France');
CALL GetOfficeByCountry(); -- Error Code: 1318. Incorrect number of arguments for PROCEDURE classicmodels.GetOfficeByCountry; expected 1, got 0

-- OUT param

-- The value of an OUT parameter can be changed inside the stored procedure and its new value is passed back to the calling program.
-- the stored procedure cannot access the initial value of the OUT parameter when it starts.

DELIMITER $$

CREATE PROCEDURE GetOrderCountByStatus (
	IN  orderStatus VARCHAR(25),
	OUT total INT
)
BEGIN
	SELECT COUNT(orderNumber)
	INTO total
	FROM orders
	WHERE status = orderStatus;
END$$

DELIMITER ;

-- find the number of orders that already shipped
CALL GetOrderCountByStatus('Shipped',@total);
SELECT @total;

-- get the number of orders that are in-process
CALL GetOrderCountByStatus('in process',@total);
SELECT @total AS  total_in_process;


-- INOUT param

-- An INOUT  parameter is a combination of IN and OUT parameters
-- It means that the calling program may pass the argument, and the stored procedure can modify the INOUT parameter, and pass the new value back to the calling program.

DELIMITER $$

CREATE PROCEDURE SetCounter(
	INOUT counter INT,
    IN inc INT
)
BEGIN
	SET counter = counter + inc;
END$$

DELIMITER ;

SET @counter = 1;
CALL SetCounter(@counter,1); -- 2
CALL SetCounter(@counter,1); -- 3
CALL SetCounter(@counter,5); -- 8
SELECT @counter; -- 8