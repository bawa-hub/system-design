-- Table aliases temporarily assign tables new names during the execution of a query.
-- AS keyword is optional
-- Similar to column aliases,

-- table_name AS alias_name;
-- table_name alias_name;


-- 1) Using table aliases for the long table name to make queries more readable
-- a_very_long_table_name.column_name
-- a_very_long_table_name AS alias
-- alias.column_name


-- 2) Using table aliases in join clauses
SELECT
	c.customer_id,
	first_name,
	amount,
	payment_date
FROM
	customer c
INNER JOIN payment p 
    ON p.customer_id = c.customer_id
ORDER BY 
   payment_date DESC;

-- 3) Using table aliases in self-join
SELECT
    e.first_name employee,
    m .first_name manager
FROM
    employee e
INNER JOIN employee m 
    ON m.employee_id = e.manager_id
ORDER BY manager;
