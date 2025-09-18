-- RENAME TABLE statement
-- RENAME TABLE old_table_name TO new_table_name;

-- we can use the RENAME TABLE statement to rename views.
-- Before we execute the RENAME TABLE statement, we must ensure that there is no active transactions or locked tables.

-- Note that you cannot use the RENAME TABLE statement to rename a temporary table, 
-- but you can use the ALTER TABLE statement to rename a temporary table.


CREATE DATABASE IF NOT EXISTS hr;

CREATE TABLE departments (
    department_id INT AUTO_INCREMENT PRIMARY KEY,
    dept_name VARCHAR(100)
);

CREATE TABLE employees (
    id int AUTO_INCREMENT primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    department_id int not null,
    FOREIGN KEY (department_id)
        REFERENCES departments (department_id)
);

INSERT INTO departments(dept_name)
VALUES('Sales'),('Markting'),('Finance'),('Accounting'),('Warehouses'),('Production');

INSERT INTO employees(first_name,last_name,department_id) 
VALUES('John','Doe',1),
		('Bush','Lily',2),
		('David','Dave',3),
		('Mary','Jane',4),
		('Jonatha','Josh',5),
		('Mateo','More',1);


-- Renaming a table referenced by a view
-- If the table that you are going to rename is referenced by a view, 
-- the view will become invalid if you rename the table, and you have to adjust the view manually.

CREATE VIEW v_employee_info as
    SELECT 
        id, first_name, last_name, dept_name
    from
        employees
            inner join
        departments USING (department_id);

RENAME TABLE employees TO people;

SELECT 
    *
FROM
    v_employee_info;
-- Error Code: 1356. View 'hr.v_employee_info' references invalid table(s) or 
-- column(s) or function(s) or definer/invoker of view lack rights to use them    

-- check the status of the v_employee_info view
CHECK TABLE v_employee_info;


-- Renaming a table that referenced by a stored procedure
-- In case the table that you are going to rename is referenced by a stored procedure, 
-- you have to manually adjust it like you did with the view.

RENAME TABLE people TO employees;

DELIMITER $$

CREATE PROCEDURE get_employee(IN p_id INT)

BEGIN
	SELECT first_name
		,last_name
		,dept_name
	FROM employees
	INNER JOIN departments using (department_id)
	WHERE id = p_id;
END $$

DELIMITER;

CALL get_employee(1);

RENAME TABLE employees TO people;

CALL get_employee(2);
-- Error Code: 1146. Table 'hr.employees' doesn't exist


-- Renaming a table that has foreign keys referenced to

RENAME TABLE departments TO depts;

DELETE FROM depts 
WHERE
    department_id = 1;

-- Error Code: 1451. Cannot delete or update a parent row: a foreign key constraint fails (`hr`.`people`, CONSTRAINT `people_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `depts` (`department_id`))


-- Renaming multiple tables
RENAME TABLE old_table_name_1 TO new_table_name_2,
             old_table_name_2 TO new_table_name_2,...


RENAME TABLE depts TO departments,
             people TO employees;



-- Renaming tables using ALTER TABLE statement
ALTER TABLE old_table_name
RENAME TO new_table_name;

-- ALTER TABLE statement can rename a temporary table while the RENAME TABLE statement cannot.


-- Renaming temporary table example

CREATE TEMPORARY TABLE lastnames
SELECT DISTINCT last_name from employees;

RENAME TABLE lastnames TO unique_lastnames;
-- Error Code: 1017. Can't find file: '.\hr\lastnames.frm' (errno: 2 - No such file or directory)

ALTER TABLE lastnames
RENAME TO unique_lastnames;