-- MySQL does not support the INTERSECT operator.
-- emulate the INTERSECT operator in MySQL using join clauses.


-- INTERSECT operator is a set operator that returns only distinct rows of two queries or more queries.

-- syntax
-- (SELECT column_list 
-- FROM table_1)
-- INTERSECT
-- (SELECT column_list
-- FROM table_2);

-- rules are same as union



-- emulation intersect in mysql
CREATE TABLE t1 (
    id INT PRIMARY KEY
);
CREATE TABLE t2 LIKE t1;
INSERT INTO t1(id) VALUES(1),(2),(3);
INSERT INTO t2(id) VALUES(2),(3),(4);

SELECT id FROM t1;
-- id
-- ----
-- 1
-- 2
-- 3

SELECT id FROM t2;
-- id
-- ---
-- 2
-- 3
-- 4

-- Emulate INTERSECT using DISTINCT and INNER JOIN clause
SELECT DISTINCT 
   id 
FROM t1
   INNER JOIN t2 USING(id);

-- id
-- ----
-- 2
-- 3

-- How it works.

--     The INNER JOIN clause returns rows from both left and right tables.
--     The DISTINCT operator removes the duplicate rows.

-- Emulate INTERSECT using IN and subquery
SELECT DISTINCT id
FROM t1
WHERE id IN (SELECT id FROM t2);
-- id
-- ----
-- 2
-- 3

-- How it works.

--     The subquery returns the first result set.
--     The outer query uses the IN operator to select only values that exist in the first result set. 
-- The DISTINCT operator ensures that only distinct values are selected.