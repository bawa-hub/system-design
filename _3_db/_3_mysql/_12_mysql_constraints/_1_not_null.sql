-- NOT NULL constraint is a column constraint that ensures values stored in a column are not NULL

-- syntax
column_name data_type NOT NULL;


CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE
);

-- Add a NOT NULL constraint to an existing column
-- a. Check the current values of the column if there is any NULL.
-- b. Update the NULL  to non-NULL if NULLs exist.
-- c. Modify the column with a NOT NULL constraint.

INSERT INTO tasks(title ,start_date, end_date)
VALUES('Learn MySQL NOT NULL constraint', '2017-02-01','2017-02-02'),
      ('Check and update NOT NULL constraint to your database', '2017-02-01',NULL);


SELECT * 
FROM tasks
WHERE end_date IS NULL;  

UPDATE tasks 
SET 
    end_date = start_date + 7
WHERE
    end_date IS NULL;


ALTER TABLE tasks 
CHANGE 
    end_date 
    end_date DATE NOT NULL;


-- Drop a NOT NULL constraint
ALTER TABLE tasks 
MODIFY 
    end_date 
    end_date DATE NOT NULL;