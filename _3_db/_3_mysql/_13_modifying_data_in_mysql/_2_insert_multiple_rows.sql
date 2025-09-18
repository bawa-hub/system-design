-- insert multiple rows
-- INSERT INTO table(c1,c2,...)
-- VALUES 
--    (v11,v12,...),
--    (v21,v22,...),
--     ...
--    (vnn,vn2,...);

-- MySQL server receives the INSERT statement whose size is bigger than max_allowed_packet, it will issue a packet too large error and terminates the connection.
SHOW VARIABLES LIKE 'max_allowed_packet';

-- set a new value for the max_allowed_packet variable
SET GLOBAL max_allowed_packet=size; --size=number of max allowed packet size in bytes.


-- creating a table
CREATE TABLE projects(
	project_id INT AUTO_INCREMENT, 
	name VARCHAR(100) NOT NULL,
	start_date DATE,
	end_date DATE,
	PRIMARY KEY(project_id)
);

INSERT INTO 
	projects(name, start_date, end_date)
VALUES
	('AI for Marketing','2019-08-01','2019-12-31'),
	('ML for Sales','2019-05-15','2019-11-20');

-- when you insert multiple rows and use the LAST_INSERT_ID() function to get the last inserted id of an AUTO_INCREMENT column, you will get the id of the first inserted row only, not the id of the last inserted row.
