-- Fetch employees who earn more than avg salary of all employees

with average_salary (avg_sal) as 
      (select cast(avg(salary) as int) from employees)
select *
from employees e, average_salary av
where e.salary > av.avg_sal;
-- this can also be done with subquery


-- Find stores whose sales are better than avg sales across all stores

