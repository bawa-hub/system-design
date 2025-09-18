-- show table command:
>> SHOW tables;
-- show table with type
>> SHOW FULL tables;

-- SHOW TABLES command provides you with an option that allows you 
-- to filter the returned tables using the LIKE operator or an expression in the WHERE clause 
SHOW TABLES LIKE pattern;
SHOW TABLES WHERE expression;
SHOW TABLES LIKE '%es';
SHOW FULL TABLES WHERE table_type = 'VIEW';