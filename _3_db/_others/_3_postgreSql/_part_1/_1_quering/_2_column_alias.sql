-- syntax
SELECT column_name AS alias_name
FROM table_name;

-- AS is optional
SELECT column_name alias_name
FROM table_name;


SELECT 
   first_name, 
   last_name AS surname
FROM customer;
-- or
SELECT 
   first_name, 
   last_name surname
FROM customer;


SELECT
    first_name || ' ' || last_name AS full_name
FROM
    customer;

-- Column aliases that contain spaces
SELECT
    first_name || ' ' || last_name "full name"
FROM
    customer;    