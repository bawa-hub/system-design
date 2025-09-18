-- MySQL does not support the MINUS operator.

-- MINUS compares the results of two queries and returns distinct rows from the result set of the first query that does not appear in the result set of the second query.

-- syntax
-- SELECT select_list1 
-- FROM table_name1
-- MINUS 
-- SELECT select_list2 
-- FROM table_name2;

-- rules are same as union and intersect

CREATE TABLE t1 (
    id INT PRIMARY KEY
);

CREATE TABLE t2 (
    id INT PRIMARY KEY
);

INSERT INTO t1 VALUES (1),(2),(3);
INSERT INTO t2 VALUES (2),(3),(4);

SELECT id FROM t1
MINUS
SELECT id FROM t2; 


-- MySQL MINUS operator emulation
-- SELECT 
--     select_list
-- FROM 
--     table1
-- LEFT JOIN table2 
--     ON join_predicate
-- WHERE 
--     table2.column_name IS NULL; 

SELECT 
    id
FROM
    t1
LEFT JOIN
    t2 USING (id)
WHERE
    t2.id IS NULL;
