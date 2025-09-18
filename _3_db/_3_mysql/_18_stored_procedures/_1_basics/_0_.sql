-- a stored procedure is a segment of declarative SQL statements stored inside the MySQL Server

DELIMITER $$

CREATE PROCEDURE GetCustomers()
BEGIN
	SELECT 
		customerName, 
		city, 
		state, 
		postalCode, 
		country
	FROM
		customers
	ORDER BY customerName;    
END$$
DELIMITER ;

CALL GetCustomers();

-- first time you invoke a stored procedure, 
-- MySQL looks up for the name in the database catalog, 
-- compiles the stored procedure’s code, 
-- place it in a memory area known as a cache, and execute the stored procedure.

-- If you invoke the same stored procedure in the same session again, 
-- MySQL just executes the stored procedure from the cache without having to recompile it.


-- stored procedures advantages