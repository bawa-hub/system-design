ALTER TABLE table
ADD [COLUMN] column_name column_definition [FIRST|AFTER existing_column];

-- MySQL allows you to add the new column as the first column of the table by specifying the FIRST keyword. 
-- It also allows you to add the new column after an existing column using the AFTER existing_column clause. 
-- If you don’t explicitly specify the position of the new column, MySQL will add it as the last column.

-- add two or more columns to a table at the same time
ALTER TABLE table
ADD [COLUMN] column_name_1 column_1_definition [FIRST|AFTER existing_column],
ADD [COLUMN] column_name_2 column_2_definition [FIRST|AFTER existing_column],
...;

-- ADD COLUMN examples
CREATE TABLE IF NOT EXISTS vendors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

ALTER TABLE vendors
ADD COLUMN phone VARCHAR(15) AFTER name;

ALTER TABLE vendors
ADD COLUMN vendor_group INT NOT NULL;

INSERT INTO vendors(name,phone,vendor_group)
VALUES('IBM','(408)-298-2987',1);

INSERT INTO vendors(name,phone,vendor_group)
VALUES('Microsoft','(408)-298-2988',1);

ALTER TABLE vendors
ADD COLUMN email VARCHAR(100) NOT NULL,
ADD COLUMN hourly_rate decimal(10,2) NOT NULL;

-- If you accidentally add a column that already exists in the table, MySQL will issue an error

ALTER TABLE vendors
ADD COLUMN vendor_group INT NOT NULL;

-- Error Code: 1060. Duplicate column name 'vendor_group'



-- check whether a column already exists in a table before adding it
-- there is no statement like ADD COLUMN IF NOT EXISTS available
-- you can get this information from the columns table of the information_schema database 

SELECT 
    IF(count(*) = 1, 'Exist','Not Exist') AS result
FROM
    information_schema.columns
WHERE
    table_schema = 'classicmodels'
        AND table_name = 'vendors'
        AND column_name = 'phone';