-- SHOW DATABASES command:
>> SHOW DATABASES;
OR, 
>> SHOW SCHEMAS;

-- select database
>> USE db_name

-- query the database that matches a specific pattern,
>> SHOW DATABASES LIKE pattern;

-- If the condition in the LIKE clause is not sufficient, 
-- you can query the database information directly from the schemata table in the information_schema database.
SELECT schema_name
FROM information_schema.schemata
WHERE schema_name LIKE '%schema' OR 
      schema_name LIKE '%s';
