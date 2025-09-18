-- syntax
DROP PROCEDURE [IF EXISTS] stored_procedure_name;


--  create a new stored procedure that returns employee and office information:
DELIMITER $$

CREATE PROCEDURE GetEmployees()
BEGIN
    SELECT 
        firstName, 
        lastName, 
        city, 
        state, 
        country
    FROM employees
    INNER JOIN offices using (officeCode);
    
END$$

DELIMITER ;

-- drop GetEmployee() stored procedure
DROP PROCEDURE GetEmployees; -- may give warning if not exists
DROP PROCEDURE IF EXISTS GetEmployee