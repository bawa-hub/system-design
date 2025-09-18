-- MySQL DROP COLUMN statement
ALTER TABLE table_name
DROP COLUMN column_name;

-- COLUMN keyword is optional
ALTER TABLE table_name
DROP column_name;

-- remove multiple columns from a table
ALTER TABLE table_name
DROP COLUMN column_name_1,
DROP COLUMN column_name_2,
...;

    -- Removing a column from a table makes all database objects such as stored procedures, views, and triggers that referencing the dropped column invalid. 
    -- For example, you may have a stored procedure that refers to a column. When you remove the column, the stored procedure becomes invalid. 
    -- To fix it, you have to manually change the stored procedure’s code.
    -- The code from other applications that depends on the dropped column must be also changed, which takes time and efforts.
    -- Dropping a column from a large table can impact the performance of the database during the removal time.

    CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    excerpt VARCHAR(400),
    content TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

ALTER TABLE posts
DROP COLUMN excerpt;

DESCRIBE posts;

ALTER TABLE posts
DROP COLUMN created_at,
DROP COLUMN updated_at;

DESCRIBE posts;


-- drop a column which is a foreign key
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

ALTER TABLE posts 
ADD COLUMN category_id INT NOT NULL;

ALTER TABLE posts 
ADD CONSTRAINT fk_cat 
FOREIGN KEY (category_id) 
REFERENCES categories(id);

ALTER TABLE posts
DROP COLUMN category_id;

-- Error Code: 1553. Cannot drop index 'fk_cat': needed in a foreign key constraint

-- To avoid this error, you must remove the foreign key constraint before dropping the column.