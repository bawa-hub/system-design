-- ALL operator allows you to query data by comparing a value with a list of values returned by a subquery.
comparison_operator ALL (subquery)


-- With the assumption that the subquery returns some rows, the ALL operator works as:

    -- column_name > ALL (subquery) the expression evaluates to true if a value is greater than the biggest value returned by the subquery.
    -- column_name >= ALL (subquery) the expression evaluates to true if a value is greater than or equal to the biggest value returned by the subquery.
    -- column_name < ALL (subquery) the expression evaluates to true if a value is less than the smallest value returned by the subquery.
    -- column_name <= ALL (subquery) the expression evaluates to true if a value is less than or equal to the smallest value returned by the subquery.
    -- column_name = ALL (subquery) the expression evaluates to true if a value is equal to any value returned by the subquery.
    -- column_name != ALL (subquery) the expression evaluates to true if a value is not equal to any value returned by the subquery.

    -- In case the subquery returns no row, then the ALL operator always evaluates to true.


-- returns the average lengths of all films grouped by film rating
SELECT
    ROUND(AVG(length), 2) avg_length
FROM
    film
GROUP BY
    rating
ORDER BY
    avg_length DESC;


-- find all films whose lengths are greater than the list of the average lengths above
SELECT
    film_id,
    title,
    length
FROM
    film
WHERE
    length > ALL (
            SELECT
                ROUND(AVG (length),2)
            FROM
                film
            GROUP BY
                rating
    )
ORDER BY
    length;
