-- comparison operators --

-- Operator	Description
-- =	Equal to. You can use it with almost any data type.
-- <> or !=	Not equal to
-- <	Less than. You typically use it with numeric and date/time data types.
-- >	Greater than.
-- <=	Less than or equal to
-- >=	Greater than or equal to

-- find all employees who are not the Sales Rep

SELECT 
    lastname, 
    firstname, 
    jobtitle
FROM
    employees
WHERE
    jobtitle <> 'Sales Rep';

-- finds employees whose office code is greater than 5

SELECT 
    lastname, 
    firstname, 
    officeCode
FROM
    employees
WHERE 
    officecode > 5;