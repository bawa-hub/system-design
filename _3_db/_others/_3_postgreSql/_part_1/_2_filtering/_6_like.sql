-- to query data using pattern matchings.

-- PostgreSQL provides you with two wildcards:
--     Percent sign ( %) matches any sequence of zero or more characters.
--     Underscore sign ( _)  matches any single character.

-- returns customers whose first name contains  er string like Jenifer, Kimberly
SELECT
	first_name,
        last_name
FROM
	customer
WHERE
	first_name LIKE '%er%'
ORDER BY 
        first_name;


-- combine the percent ( %) with underscore ( _) to construct a pattern
SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	first_name LIKE '_her%'
ORDER BY 
        first_name;


-- find customers whose first names do not begin with Jen
SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	first_name NOT LIKE 'Jen%'
ORDER BY 
        first_name



-- PostgreSQL supports the ILIKE operator that works like the LIKE operator. 
-- In addition, the ILIKE operator matches value case-insensitively

SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	first_name ILIKE 'BAR%';


-- PostgreSQL also provides some operators that act like the LIKE, NOT LIKE, ILIKE and NOT ILIKE operator
-- ~~	LIKE
-- ~~*	ILIKE
-- !~~	NOT LIKE
-- !~~*	NOT ILIKE