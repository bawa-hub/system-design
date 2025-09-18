-- find the films whose rental rate is higher than the average rental rate
SELECT -- outer query/main query
	film_id,
	title,
	rental_rate
FROM
	film
WHERE
	rental_rate > (
		SELECT -- inner query/subquery
			AVG (rental_rate)
		FROM
			film
	);

-- PostgreSQL executes the query that contains a subquery in the following sequence:
--     First, executes the subquery.
--     Second, gets the result and passes it to the outer query.
--     Third, executes the outer query.


-- PostgreSQL subquery with IN operator


-- films that have the returned date between 2005-05-29 and 2005-05-30
SELECT
	film_id,
	title
FROM
	film
WHERE
	film_id IN (
		SELECT
			inventory.film_id
		FROM
			rental
		INNER JOIN inventory ON inventory.inventory_id = rental.inventory_id
		WHERE
			return_date BETWEEN '2005-05-29'
		AND '2005-05-30'
	);


-- PostgreSQL subquery with EXISTS operator


-- If the subquery returns any row, the EXISTS operator returns true. 
-- If the subquery returns no row, the result of EXISTS operator is false.

-- EXISTS operator only cares about the number of rows returned from the subquery, not the content of the rows, 
-- therefore, the common coding convention of EXISTS operator is:
EXISTS (SELECT 1 FROM tbl WHERE condition);

SELECT
	first_name,
	last_name
FROM
	customer
WHERE
	EXISTS (
		SELECT
			1
		FROM
			payment
		WHERE
			payment.customer_id = customer.customer_id
	);