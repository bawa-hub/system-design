-- ANY operator compares a value to a set of values returned by a subquery
expresion operator ANY(subquery)

    -- The subquery must return exactly one column.
    -- The ANY operator must be preceded by one of the following comparison operator =, <=, >, <, > and <>
    -- The ANY operator returns true if any value of the subquery meets the condition, otherwise, it returns false.

    -- SOME is a synonym for ANY, meaning that you can substitute SOME for ANY in any SQL statement.


-- returns the maximum length of film grouped by film category
SELECT
    MAX( length )
FROM
    film
INNER JOIN film_category
        USING(film_id)
GROUP BY
    category_id;


-- finds the films whose lengths are greater than or equal to the maximum length of any film category
SELECT title
FROM film
WHERE length >= ANY(
    SELECT MAX( length )
    FROM film
    INNER JOIN film_category USING(film_id)
    GROUP BY  category_id );



-- ANY vs. IN
-- The = ANY is equivalent to IN operator.


-- film whose category is either Action or Drama
SELECT
    title,
    category_id
FROM
    film
INNER JOIN film_category
        USING(film_id)
WHERE
    category_id = ANY(
        SELECT
            category_id
        FROM
            category
        WHERE
            NAME = 'Action'
            OR NAME = 'Drama'
    );

-- equivalent to

SELECT
    title,
    category_id
FROM
    film
INNER JOIN film_category
        USING(film_id)
WHERE
    category_id IN(
        SELECT
            category_id
        FROM
            category
        WHERE
            NAME = 'Action'
            OR NAME = 'Drama'
    );



-- <> ANY operator is different from NOT IN
x <> ANY (a,b,c) 
-- equivalent to
x <> a OR <> b OR x <> c