-- recursive common table expression (CTE) is a CTE that has a subquery which refers to the CTE name itself. 

-- syntax
WITH RECURSIVE cte_name AS (
    initial_query  -- anchor member
    UNION ALL
    recursive_query -- recursive member that references to the CTE name
)
SELECT * FROM cte_name;


-- recursive CTE consists of three main parts:

-- a. An initial query that forms the base result set of the CTE structure. 
-- The initial query part is referred to as an anchor member.

-- b. A recursive query part is a query that references to the CTE name, 
-- therefore, it is called a recursive member. 
-- The recursive member is joined with the anchor member by aUNION ALL or UNION DISTINCT operator.

-- c. A termination condition that ensures the recursion stops when the recursive member returns no row.

-- Refer: https://www.mysqltutorial.org/mysql-recursive-cte/