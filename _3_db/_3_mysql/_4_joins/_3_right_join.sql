-- right join

SELECT 
    select_list
FROM t1
RIGHT JOIN t2 ON 
    join_condition;


-- 1) join the table customers with the table employees using right join
SELECT 
    employeeNumber, 
    customerNumber
FROM
    customers
RIGHT JOIN employees 
    ON salesRepEmployeeNumber = employeeNumber
ORDER BY 
	employeeNumber;


-- 2) Using MySQL RIGHT JOIN to find unmatching rows
-- find employees who are not in charge of any customers:
SELECT 
    employeeNumber, 
    customerNumber
FROM
    customers
RIGHT JOIN employees ON 
	salesRepEmployeeNumber = employeeNumber
WHERE customerNumber is NULL
ORDER BY employeeNumber;