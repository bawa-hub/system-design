-- To create a sequence in MySQL automatically, you set the AUTO_INCREMENT attribute for a column, 
-- which typically is a primary key column

-- following rules are applied when you use the AUTO_INCREMENT attribute:
--     Each table has only one AUTO_INCREMENT column whose data type is typically the integer.
--     The  AUTO_INCREMENT column must be indexed, which means it can be either PRIMARY KEY or UNIQUE index.
--     The AUTO_INCREMENT column must have a NOT NULL constraint. When you set the AUTO_INCREMENT attribute to a column, MySQL automatically adds the NOT NULL  constraint to the column implicitly.

CREATE TABLE employees (
    emp_no INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50)
);

-- https://www.mysqltutorial.org/mysql-sequence/