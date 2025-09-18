-- create procedure to returns all products in the products table

DELIMITER //

CREATE PROCEDURE GetAllProducts()
BEGIN
	SELECT *  FROM products;
END //

DELIMITER ;

-- CALL THE PROCEDURE
CALL GetAllProducts();