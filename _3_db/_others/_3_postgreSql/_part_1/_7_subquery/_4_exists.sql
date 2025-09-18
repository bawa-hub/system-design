-- EXISTS operator is a boolean operator that tests for existence of rows in a subquery
EXISTS (subquery)

-- If the subquery returns at least one row, the result of EXISTS is true. 
-- In case the subquery returns no row, the result is of EXISTS is false.

-- The EXISTS operator is often used with the correlated subquery.



-- Find customers who have at least one payment whose amount is greater than 11.
SELECT first_name,
       last_name
FROM customer c
WHERE EXISTS
    (SELECT 1
     FROM payment p
     WHERE p.customer_id = c.customer_id
       AND amount > 11 )
ORDER BY first_name,
         last_name;


-- returns customers have not made any payment that greater than 11.
SELECT first_name,
       last_name
FROM customer c
WHERE NOT EXISTS
    (SELECT 1
     FROM payment p
     WHERE p.customer_id = c.customer_id
       AND amount > 11 )
ORDER BY first_name,
         last_name;



-- If the subquery returns NULL, EXISTS returns true.

SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	EXISTS( SELECT NULL )
ORDER BY
	first_name,
	last_name;
