-- insert a row
-- INSERT INTO table(c1,c2,...)
-- VALUES (v1,v2,...);

-- insert multiple rows
-- INSERT INTO table(c1,c2,...)
-- VALUES 
--    (v11,v12,...),
--    (v21,v22,...),
--     ...
--    (vnn,vn2,...);

-- creating a table
CREATE TABLE IF NOT EXISTS tasks (
    task_id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    due_date DATE,
    priority TINYINT NOT NULL DEFAULT 3,
    description TEXT,
    PRIMARY KEY (task_id)
);

-- simple INSERT example
INSERT INTO tasks(title,priority)
VALUES('Learn MySQL INSERT Statement',1);

-- inserting rows using default value
INSERT INTO tasks(title,priority)
VALUES('Understanding DEFAULT keyword in INSERT statement',DEFAULT);

-- inserting dates
INSERT INTO tasks(title, start_date, due_date)
VALUES('Insert date into table','2018-01-09','2018-09-15');

-- using expression in values
INSERT INTO tasks(title,start_date,due_date)
VALUES('Use current date for the task',CURRENT_DATE(),CURRENT_DATE())

-- inserting multiple values
INSERT INTO tasks(title, priority)
VALUES
	('My first task', 1),
	('It is the second task',2),
	('This is the third task of the week',3);