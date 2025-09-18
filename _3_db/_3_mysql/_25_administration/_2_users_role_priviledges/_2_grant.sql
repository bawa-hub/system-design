-- syntax
-- GRANT privilege [,privilege],.. 
-- ON privilege_level 
-- TO account_name;


-- grants the SELECT privilege on the table employees
GRANT SELECT
ON employees
TO bawa@localhost;

-- grants UPDATE, DELETE, and INSERT privileges on the table employees
GRANT INSERT, UPDATE, DELETE
ON employees 
TO bawa@localhost;


-- Privilege Levels
Global,
Database,
Table,
Column,
Stored Routine,
Proxy


-- https://www.mysqltutorial.org/mysql-grant.aspx