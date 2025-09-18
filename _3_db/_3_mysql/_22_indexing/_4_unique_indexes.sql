-- MySQL UNIQUE index to prevent duplicate values in one or more columns in a table.
CREATE UNIQUE INDEX index_name
ON table_name(index_column_1,index_column_2,...);

-- when you create a table.
CREATE TABLE table_name(
...
   UNIQUE KEY(index_column_,index_column_2,...) 
);

or

CREATE TABLE table_name(
...
   UNIQUE INDEX(index_column_,index_column_2,...) 
);

-- add a unique constraint to an existing table
ALTER TABLE table_name
ADD CONSTRAINT constraint_name UNIQUE KEY(column_1,column_2,...);


-- MySQL UNIQUE Index & NULL

-- Unlike other database systems, MySQL considers NULL values as distinct values. 
-- Therefore, you can have multiple NULL values in the UNIQUE index.
-- UNIQUE constraint does not apply to NULL values except for the BDB storage engine.

CREATE TABLE IF NOT EXISTS contacts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    email VARCHAR(100) NOT NULL,
    UNIQUE KEY unique_email (email)
);

--  combination
CREATE UNIQUE INDEX idx_name_phone
ON contacts(first_name,last_name,phone);