-- A CROSS JOIN clause allows you to produce a Cartesian Product of rows in two or more tables.

-- syntax

SELECT select_list
FROM T1
CROSS JOIN T2;

-- or

SELECT select_list
FROM T1, T2;

-- or

SELECT *
FROM T1
INNER JOIN T2 ON true;


-- sample data for demonstration
DROP TABLE IF EXISTS T1;
CREATE TABLE T1 (label CHAR(1) PRIMARY KEY);

DROP TABLE IF EXISTS T2;
CREATE TABLE T2 (score INT PRIMARY KEY);

INSERT INTO T1 (label)
VALUES
	('A'),
	('B');

INSERT INTO T2 (score)
VALUES
	(1),
	(2),
	(3);


SELECT *
FROM T1
CROSS JOIN T2;