-- MySQL treats the views as tables with the type 'VIEW'

SHOW FULL TABLES 
WHERE table_type = 'VIEW';

-- FROM ANTORHER DATABASE
SHOW FULL TABLES
[{FROM | IN } database_name]
WHERE table_type = 'VIEW';

SHOW FULL TABLES IN sys 
WHERE table_type='VIEW';


-- get views that match a pattern,
SHOW FULL TABLES 
FROM sys
LIKE 'waits%';


-- MySQL Show View –  using INFORMATION_SCHEMA database

-- INFORMATION_SCHEMA database provides access to MySQL database metadata such as databases, tables, data types of columns, or privileges.
-- INFORMATION_SCHEMA is sometimes referred to as a database dictionary or system catalog.

SELECT * 
FROM information_schema.tables;

-- columns which are relevant to the views are:
--     The table_schema column stores the schema or database of the view (or table).
--     The table_name column stores the name of the view (or table).
--     The table_type column stores the type of tables: BASE TABLE for a table, VIEW for a view, or SYSTEM VIEW for an INFORMATION_SCHEMA table.

-- returns all views from the classicmodels database:
SELECT 
    table_name view_name
FROM 
    information_schema.tables 
WHERE 
    table_type   = 'VIEW' AND 
    table_schema = 'classicmodels';

-- finds all views whose names start with customer
SELECT 
    table_name view_name
FROM 
    information_schema.tables 
WHERE 
    table_type   = 'VIEW' AND 
    table_schema = 'classicmodels' AND
    table_name   LIKE 'customer%';
