CREATE TABLE vehicles (
    vehicleId INT,
    year INT NOT NULL,
    make VARCHAR(100) NOT NULL,
    PRIMARY KEY(vehicleId)
);

-- Add a column to a table
-- ALTER TABLE table_name
-- ADD 
--     new_column_name column_definition
--     [FIRST | AFTER column_name]

ALTER TABLE vehicles
ADD model VARCHAR(100) NOT NULL;

DESCRIBE vehicles;

-- Add multiple columns to a table
-- ALTER TABLE table_name
--     ADD new_column_name column_definition
--     [FIRST | AFTER column_name],
--     ADD new_column_name column_definition
--     [FIRST | AFTER column_name],
--     ...;

ALTER TABLE vehicles
ADD color VARCHAR(50),
ADD note VARCHAR(255);

-- Modify a column
-- ALTER TABLE table_name
-- MODIFY column_name column_definition
-- [ FIRST | AFTER column_name];    

ALTER TABLE vehicles 
MODIFY note VARCHAR(100) NOT NULL;

ALTER TABLE vehicles 
MODIFY year SMALLINT NOT NULL,
MODIFY color VARCHAR(20) NULL AFTER make;

-- Rename a column in a table
-- ALTER TABLE table_name
--     CHANGE COLUMN original_name new_name column_definition
--     [FIRST | AFTER column_name];

ALTER TABLE vehicles 
CHANGE COLUMN note vehicleCondition VARCHAR(100) NOT NULL;

-- Drop a column
-- ALTER TABLE table_name
-- DROP COLUMN column_name;

ALTER TABLE vehicles
DROP COLUMN vehicleCondition;

-- Rename table
-- ALTER TABLE table_name
-- RENAME TO new_table_name;

ALTER TABLE vehicles 
RENAME TO cars; 
