-- get the nth highest or lowest value

SELECT select_list
FROM table_name
ORDER BY sort_expression
LIMIT n-1, 1;

-- find the customer who has the second-highest credit:

SELECT 
    customerName, 
    creditLimit
FROM
    customers
ORDER BY 
    creditLimit DESC    
LIMIT 1,1;

-- Note that this technique works when there are no two customers who have the same credit limits. 
-- To get a more accurate result, you should use the DENSE_RANK() window function.

-- return the first five unique states in the customers table:

SELECT DISTINCT
    state
FROM
    customers
WHERE
    state IS NOT NULL
LIMIT 5;