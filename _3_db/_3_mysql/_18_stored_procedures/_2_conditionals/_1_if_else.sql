-- IF-THEN statement
-- IF condition THEN 
--    statements;
-- END IF;

DELIMITER $$

CREATE PROCEDURE GetCustomerLevel(
    IN  pCustomerNumber INT, 
    OUT pCustomerLevel  VARCHAR(20))
BEGIN
    DECLARE credit DECIMAL(10,2) DEFAULT 0;

    SELECT creditLimit 
    INTO credit
    FROM customers
    WHERE customerNumber = pCustomerNumber;

    IF credit > 50000 THEN
        SET pCustomerLevel = 'PLATINUM';
    END IF;
END$$

DELIMITER ;


-- finds all customers that have a credit limit greater than 50,000
SELECT 
    customerNumber, 
    creditLimit
FROM 
    customers
WHERE 
    creditLimit > 50000
ORDER BY 
    creditLimit DESC;


CALL GetCustomerLevel(141, @level);
SELECT @level;


-- IF-THEN-ELSE statement
-- IF condition THEN
--    statements;
-- ELSE
--    else-statements;
-- END IF;

DROP PROCEDURE GetCustomerLevel;

DELIMITER $$

CREATE PROCEDURE GetCustomerLevel(
    IN  pCustomerNumber INT, 
    OUT pCustomerLevel  VARCHAR(20))
BEGIN
    DECLARE credit DECIMAL DEFAULT 0;

    SELECT creditLimit 
    INTO credit
    FROM customers
    WHERE customerNumber = pCustomerNumber;

    IF credit > 50000 THEN
        SET pCustomerLevel = 'PLATINUM';
    ELSE
        SET pCustomerLevel = 'NOT PLATINUM';
    END IF;
END$$

DELIMITER ;

SELECT 
    customerNumber, 
    creditLimit
FROM 
    customers
WHERE 
    creditLimit <= 50000
ORDER BY 
    creditLimit DESC;


CALL GetCustomerLevel(447, @level);
SELECT @level;


-- IF-THEN-ELSEIF-ELSE statement
-- IF condition THEN
--    statements;
-- ELSEIF elseif-condition THEN
--    elseif-statements;
-- ...
-- ELSE
--    else-statements;
-- END IF;

DROP PROCEDURE GetCustomerLevel;

DELIMITER $$

CREATE PROCEDURE GetCustomerLevel(
    IN  pCustomerNumber INT, 
    OUT pCustomerLevel  VARCHAR(20))
BEGIN
    DECLARE credit DECIMAL DEFAULT 0;

    SELECT creditLimit 
    INTO credit
    FROM customers
    WHERE customerNumber = pCustomerNumber;

    IF credit > 50000 THEN
        SET pCustomerLevel = 'PLATINUM';
    ELSEIF credit <= 50000 AND credit > 10000 THEN
        SET pCustomerLevel = 'GOLD';
    ELSE
        SET pCustomerLevel = 'SILVER';
    END IF;
END $$

DELIMITER ;


CALL GetCustomerLevel(447, @level); 
SELECT @level;
