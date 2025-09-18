-- where inserting multiple rows INSERT IGNORE statement, the rows with invalid data that cause the error are ignored and the rows with valid data are inserted into the table.
-- INSERT IGNORE INTO table(column_list)
-- VALUES( value_list),
--       ( value_list),
--       ...

CREATE TABLE subscribers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO subscribers(email)
VALUES('john.doe@gmail.com');

-- inseerting with duplicate entries
INSERT INTO subscribers(email)
VALUES('john.doe@gmail.com'), 
      ('jane.smith@ibm.com');

-- Error Code: 1062. Duplicate entry 'john.doe@gmail.com' for key 'email'

INSERT IGNORE INTO subscribers(email)
VALUES('john.doe@gmail.com'), 
      ('jane.smith@ibm.com');
-- 1 row(s) affected, 1 warning(s): 1062 Duplicate entry 'john.doe@gmail.com' for key 'email' Records: 2  Duplicates: 1  Warnings: 1

SHOW WARNINGS;



-- MySQL INSERT IGNORE and STRICT mode

-- if you use the INSERT IGNORE statement, MySQL will issue a warning instead of an error. In addition, it will try to adjust the values to make them valid before adding the value to the table.

CREATE TABLE tokens (
    s VARCHAR(6)
);


INSERT INTO tokens VALUES('abcdefg');

-- MySQL issued the following error because the strict mode is on.
-- Error Code: 1406. Data too long for column 's' at row 1

INSERT IGNORE INTO tokens VALUES('abcdefg');
