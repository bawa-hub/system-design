-- get a unique combination of city and state from the customers table

SELECT DISTINCT
    state, city
FROM
    customers
WHERE
    state IS NOT NULL
ORDER BY 
    state, 
    city;

-- find customers who locate in California (CA), USA

SELECT 
    customername, 
    country, 
    state
FROM
    customers
WHERE
    country = 'USA' AND 
    state = 'CA';


-- customers who locate in California, USA, and have a credit limit greater than 100K.

SELECT 
    customername, 
    country, 
    state, 
    creditlimit
FROM
    customers
WHERE
    country = 'USA' AND 
    state = 'CA' AND 
    creditlimit > 100000;


--  select the customers who locate in the USA or France and have a credit limit greater than 100,000.

SELECT   
	customername, 
	country, 
	creditLimit
FROM   
	customers
WHERE(country = 'USA'
		OR country = 'France')
	  AND creditlimit > 100000;

-- return the customers who locate in the USA or the customers located in France with a credit limit greater than 100,000.      

SELECT    
    customername, 
    country, 
    creditLimit
FROM    
    customers
WHERE 
    country = 'USA'
    OR country = 'France'
    AND creditlimit > 100000;