-- to retrieve a portion of rows returned by a query.

-- syntax
OFFSET start { ROW | ROWS }
FETCH { FIRST | NEXT } [ row_count ] { ROW | ROWS } ONLY


-- FETCH vs. LIMIT
-- FETCH clause is functionally equivalent to the LIMIT clause. 
-- If you plan to make your application compatible with other database systems,
-- you should use the FETCH clause because it follows the standard SQL

-- select the first film sorted by titles in ascending order:
SELECT
    film_id,
    title
FROM
    film
ORDER BY
    title 
FETCH FIRST ROW ONLY; -- equivalent to  FETCH FIRST 1 ROW ONLY;

-- select the first five films sorted by titles
SELECT
    film_id,
    title
FROM
    film
ORDER BY
    title 
FETCH FIRST 5 ROW ONLY;


-- returns the next five films after the first five films sorted by titles
SELECT
    film_id,
    title
FROM
    film
ORDER BY
    title 
OFFSET 5 ROWS 
FETCH FIRST 5 ROW ONLY; 
