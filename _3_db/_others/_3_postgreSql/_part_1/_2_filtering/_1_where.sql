-- syntax
SELECT select_list
FROM table_name
WHERE condition
ORDER BY sort_expression

-- comparison and logical operators
-- =	Equal
-- >	Greater than
-- <	Less than
-- >=	Greater than or equal
-- <=	Less than or equal
-- <> or !=	Not equal
-- AND	Logical operator AND
-- OR	Logical operator OR
-- IN	Return true if a value matches any value in a list
-- BETWEEN	Return true if a value is between a range of values
-- LIKE	Return true if a value matches a pattern
-- IS NULL	Return true if a value is NULL
-- NOT	Negate the result of other operators


-- customers whose first names are Jamie
SELECT
	last_name,
	first_name
FROM
	customer
WHERE
	first_name = 'Jamie';


-- finds customers whose first name and last name are Jamie and rice
SELECT
	last_name,
	first_name
FROM
	customer
WHERE
	first_name = 'Jamie' AND 
        last_name = 'Rice';    


-- finds the customers whose last name is Rodriguez or first name is Adam 
SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	last_name = 'Rodriguez' OR 
	first_name = 'Adam';


-- whose first name is Ann, or Anne, or Annie
SELECT
	first_name,
	last_name
FROM
	customer
WHERE 
	first_name IN ('Ann','Anne','Annie');


-- returns all customers whose first names start with the string Ann
SELECT
	first_name,
	last_name
FROM
	customer
WHERE 
	first_name LIKE 'Ann%'


-- finds customers whose first names start with the letter A and contains 3 to 5 characters
SELECT
	first_name,
	LENGTH(first_name) name_length
FROM
	customer
WHERE 
	first_name LIKE 'A%' AND
	LENGTH(first_name) BETWEEN 3 AND 5
ORDER BY
	name_length;


-- finds customers whose first names start with Bra and last names are not Motley
SELECT 
	first_name, 
	last_name
FROM 
	customer 
WHERE 
	first_name LIKE 'Bra%' AND 
	last_name <> 'Motley';