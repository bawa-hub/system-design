-- TRUNCATE TABLE statement allows you to delete all data in a table.

-- TRUNCATE TABLE statement is like a DELETE statement without a WHERE clause that deletes all rows from a table, 
-- or a sequence of DROP TABLE and CREATE TABLE statements.

-- TRUNCATE TABLE statement is more efficient than the DELETE statement because it drops and recreates the table instead of deleting rows one by one.
TRUNCATE [TABLE] table_name;

-- Because a truncate operation causes an implicit commit, therefore, it cannot be rolled back.
-- TRUNCATE TABLE statement resets value in the AUTO_INCREMENT  column to its start value if the table has an AUTO_INCREMENT column.


-- TRUNCATE TABLE example

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL
)  ENGINE=INNODB;


DELIMITER $$
CREATE PROCEDURE load_book_data(IN num INT(4))
BEGIN
	DECLARE counter INT(4) DEFAULT 0;
	DECLARE book_title VARCHAR(255) DEFAULT '';

	WHILE counter < num DO
	  SET book_title = CONCAT('Book title #',counter);
	  SET counter = counter + 1;

	  INSERT INTO books(title)
	  VALUES(book_title);
	END WHILE;
END$$

DELIMITER ;

CALL load_book_data(10000);

TRUNCATE TABLE books;