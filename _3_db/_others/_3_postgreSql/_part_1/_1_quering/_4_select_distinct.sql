-- DISTINCT clause is used in the SELECT statement to remove duplicate rows from a result set

-- syntax
SELECT
   DISTINCT column1
FROM
   table_name;


-- PostgreSQL also provides the DISTINCT ON (expression) to keep the “first” row of each group of duplicates
SELECT
   DISTINCT ON (column1) column_alias,
   column2
FROM
   table_name
ORDER BY
   column1,
   column2;

-- use the ORDER BY clause with the DISTINCT ON(expression) to make the result set predictable

-- DISTINCT ON expression must match the leftmost expression in the ORDER BY clause.


-- data for demonstration
CREATE TABLE distinct_demo (
	id serial NOT NULL PRIMARY KEY,
	bcolor VARCHAR,
	fcolor VARCHAR
);
INSERT INTO distinct_demo (bcolor, fcolor)
VALUES
	('red', 'red'),
	('red', 'red'),
	('red', NULL),
	(NULL, 'red'),
	('red', 'green'),
	('red', 'blue'),
	('green', 'red'),
	('green', 'blue'),
	('green', 'green'),
	('blue', 'red'),
	('blue', 'green'),
	('blue', 'blue');

SELECT
	id,
	bcolor,
	fcolor
FROM
	distinct_demo ;


-- selects unique values in the  bcolor column from the t1 table and sorts the result set in alphabetical order
SELECT
	DISTINCT bcolor
FROM
	distinct_demo
ORDER BY
	bcolor;


-- DISTINCT clause on multiple columns
SELECT
	DISTINCT bcolor,
	fcolor
FROM
	distinct_demo
ORDER BY
	bcolor,
	fcolor;


-- sorts the result set by the  bcolor and  fcolor, and then for each group of duplicates, 
-- keeps the first row in the returned result set.

SELECT
	DISTINCT ON (bcolor) bcolor,
	fcolor
FROM
	distinct_demo 
ORDER BY
	bcolor,
	fcolor;