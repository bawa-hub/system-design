-- syntax
SELECT
	select_list
FROM
	table_name
ORDER BY
	sort_expression1 [ASC | DESC],
        ...
	sort_expressionN [ASC | DESC];


-- 1) Using PostgreSQL ORDER BY clause to sort rows by one column
SELECT
	first_name,
	last_name
FROM
	customer
ORDER BY
	first_name ASC; -- ASC option is the default

-- 2) Using PostgreSQL ORDER BY clause to sort rows by one column in descending order
SELECT
       first_name,
       last_name
FROM
       customer
ORDER BY
       last_name DESC;

-- 3) Using PostgreSQL ORDER BY clause to sort rows by multiple columns
SELECT
	first_name,
	last_name
FROM
	customer
ORDER BY
	first_name ASC,
	last_name DESC;

-- 4) Using PostgreSQL ORDER BY clause to sort rows by expressions
SELECT 
	first_name,
	LENGTH(first_name) len
FROM
	customer
ORDER BY 
	len DESC;



-- PostgreSQL ORDER BY clause and NULL
ORDER BY sort_expresssion [ASC | DESC] [NULLS FIRST | NULLS LAST]
-- The NULLS FIRST option places NULL before other non-null values and the NULL LAST option places NULL after other non-null values.



-- create data from demonstration

CREATE TABLE sort_demo(
	num INT
);
INSERT INTO sort_demo(num)
VALUES(1),(2),(3),(null);

SELECT num
FROM sort_demo
ORDER BY num;

-- if you use the ASC option, the ORDER BY clause uses the NULLS LAST option by default
