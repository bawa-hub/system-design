--  joins a table to itself using the inner join or left join.
-- The self join is often used to query hierarchical data or to compare a row with other rows within the same table.

-- 1) MySQL self join using INNER JOIN clause

-- get the whole organization structure
SELECT 
    CONCAT(m.lastName, ', ', m.firstName) AS Manager,
    CONCAT(e.lastName, ', ', e.firstName) AS 'Direct report'
FROM
    employees e
INNER JOIN employees m ON 
    m.employeeNumber = e.reportsTo
ORDER BY 
    Manager;

-- 2) MySQL self join using LEFT JOIN clause    

-- include the President:
SELECT 
    IFNULL(CONCAT(m.lastname, ', ', m.firstname),
            'Top Manager') AS 'Manager',
    CONCAT(e.lastname, ', ', e.firstname) AS 'Direct report'
FROM
    employees e
LEFT JOIN employees m ON 
    m.employeeNumber = e.reportsto
ORDER BY 
    manager DESC;


-- 3) Using MySQL self join to compare successive rows

-- display a list of customers who locate in the same city by joining the customers table to itself.
SELECT 
    c1.city, 
    c1.customerName, 
    c2.customerName
FROM
    customers c1
INNER JOIN customers c2 ON 
    c1.city = c2.city
    AND c1.customername > c2.customerName
ORDER BY 
    c1.city;