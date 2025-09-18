-- syntax
-- SHOW PROCEDURE STATUS [LIKE 'pattern' | WHERE search_condition]

-- shows all stored procedure in the current MySQL server
SHOW PROCEDURE STATUS;
-- lists all stored procedures in the sample database classicmodels:
SHOW PROCEDURE STATUS WHERE db = 'classicmodels';

-- ind stored procedures whose names contain a specific word,
SHOW PROCEDURE STATUS LIKE '%pattern%'
-- shows all stored procedure whose names contain the wordOrder
SHOW PROCEDURE STATUS LIKE '%Order%'


-- Listing stored procedures using the data dictionary
-- The routines table in the information_schema database contains all information on the stored procedures and stored functions of all databases in the current MySQL server.

--  show all stored procedures of a particular database
-- SELECT 
--     routine_name
-- FROM
--     information_schema.routines
-- WHERE
--     routine_type = 'PROCEDURE'
--         AND routine_schema = '<database_name>';

-- lists all stored procedures in the classicmodels database:        
SELECT 
    routine_name
FROM
    information_schema.routines
WHERE
    routine_type = 'PROCEDURE'
        AND routine_schema = 'classicmodels';