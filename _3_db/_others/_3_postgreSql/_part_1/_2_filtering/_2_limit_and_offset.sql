-- syntax
SELECT select_list 
FROM table_name
ORDER BY sort_expression
LIMIT row_count

-- In case you want to skip a number of rows before returning the row_count rows, you use OFFSET clause placed after the LIMIT clause
SELECT select_list
FROM table_name
LIMIT row_count OFFSET row_to_skip;



-- get the first five films sorted by film_id
SELECT
	film_id,
	title,
	release_year
FROM
	film
ORDER BY
	film_id
LIMIT 5;


-- retrieve 4 films starting from the fourth one ordered by film_id
SELECT
	film_id,
	title,
	release_year
FROM
	film
ORDER BY
	film_id
LIMIT 4 OFFSET 3;


-- get the top 10 most expensive films in terms of rental,
SELECT
	film_id,
	title,
	rental_rate
FROM
	film
ORDER BY
	rental_rate DESC
LIMIT 10;